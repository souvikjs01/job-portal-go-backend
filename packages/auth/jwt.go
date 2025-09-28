package auth

import (
	"fmt"
	"job_portal/packages/config"
	"job_portal/packages/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
	ValidateAccessToken(tokenString string) (*models.JWTClaims, error)
	HashPassword(password string) (string, error)
	ValidatePassword(password, hashedPassword string) error
}

type jwtService struct {
	config *config.AuthConfig
}

type CustomClaims struct {
	UserID   string      `json:"userId"`
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTService(config *config.AuthConfig) JWTService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) GenerateToken(user *models.User) (string, error) {
	// Generate access token
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}

	return accessToken, nil
}

func (s *jwtService) generateAccessToken(user *models.User) (string, error) {
	claims := CustomClaims{
		UserID:   user.ID.String(),
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "user-service",
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.Jwt_secret))
}

func (s *jwtService) ValidateAccessToken(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.Jwt_secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &models.JWTClaims{
		UserID:   claims.UserID,
		Email:    claims.Email,
		Username: claims.Username,
		Role:     claims.Role,
	}, nil
}

func (s *jwtService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *jwtService) ValidatePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
