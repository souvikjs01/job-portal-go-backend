package services

import (
	"fmt"
	"job_portal/packages/auth"
	"job_portal/packages/models"
	"job_portal/packages/repository"
)

type UserService interface {
	Register(req *models.CreateUser) (*models.User, string, error)
	// Login(req *models.LoginRequest) (*models.User, *models.TokenPair, error)
	// Logout(token string) error
	// GetProfile(token string) (*models.User, error)
	// UpdateProfile(token string, req *models.UpdateUserRequest) (*models.User, error)
	// ChangePassword(token string, req *models.ChangePasswordRequest) error
	// DeleteProfile(token, password string) error
	// ListUsers(token string, filter models.ListUsersFilter) (*models.ListUsersResponse, error)
	// GetUser(token string, userID uuid.UUID) (*models.User, error)
	// UpdateUserRole(token string, userID uuid.UUID, role models.Role) error
	// DeactivateUser(token string, userID uuid.UUID) error
}

type userService struct {
	userRepo   repository.UserRepository
	JWTService auth.JWTService
}

func NewUserService(userRepo repository.UserRepository, jwtService auth.JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		JWTService: jwtService,
	}
}

// Authentication methods
func (s *userService) Register(req *models.CreateUser) (*models.User, string, error) {
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, "", fmt.Errorf("user with email %s already exists", req.Email)
	}

	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return nil, "", fmt.Errorf("user with username %s already exists", req.Username)
	}

	// Hash password
	hashedPassword, err := s.JWTService.HashPassword(req.Password)
	if err != nil {
		return nil, "", fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	req.Password = hashedPassword

	user, err := s.userRepo.Create(req)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token
	token, err := s.JWTService.GenerateToken(user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	return user, token, nil
}
