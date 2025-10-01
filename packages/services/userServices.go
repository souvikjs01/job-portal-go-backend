package services

import (
	"fmt"
	"job_portal/packages/auth"
	"job_portal/packages/models"
	"job_portal/packages/repository"
)

type UserService interface {
	Register(req *models.CreateUser) (*models.User, string, error)
	Login(req *models.LoginUser) (*models.User, string, error)
	GetProfile(id string) (*models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateProfile(id string, req *models.UpdateUser) (*models.User, error)
	UpdateUserRole(userID string, role models.UpdateRoleRequest) error
	// ChangePassword(token string, req *models.ChangePasswordRequest) erro
	DeleteUser(userID string) error
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

// Authentication methods
func (s *userService) Login(req *models.LoginUser) (*models.User, string, error) {

	existUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, "", fmt.Errorf("user with email %s not exists", req.Email)
	}

	err = s.JWTService.ValidatePassword(req.Password, existUser.Password)
	if err != nil {
		return nil, "", fmt.Errorf("invalid password: %w", err)
	}

	// Generate token
	token, err := s.JWTService.GenerateToken(existUser)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	return existUser, token, nil
}

func (s *userService) GetProfile(id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %v", err)
	}

	return user, nil
}

func (s *userService) GetAllUser() ([]models.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) UpdateProfile(id string, req *models.UpdateUser) (*models.User, error) {
	updatedUser, err := s.userRepo.Update(id, req)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *userService) UpdateUserRole(userID string, role models.UpdateRoleRequest) error {
	err := s.userRepo.UpdateRole(userID, role)

	fmt.Println("not updated here")
	if err != nil {
		return fmt.Errorf("failed updating role")
	}
	fmt.Println("updated ss here")
	return nil
}

func (s *userService) DeleteUser(userID string) error {
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return fmt.Errorf("user not exist")
	}

	// user exist:
	err = s.userRepo.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
