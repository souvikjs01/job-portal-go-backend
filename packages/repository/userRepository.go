package repository

import (
	"fmt"
	"job_portal/packages/models"
	"job_portal/packages/store"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.CreateUser) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(id string, user *models.UpdateUser) (*models.User, error)
	Delete(id string) error
	UpdateRole(userID string, role models.UpdateRoleRequest) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *store.DB) UserRepository {
	return &userRepository{
		db: db.DB,
	}
}

func (r *userRepository) Create(user *models.CreateUser) (*models.User, error) {
	newUser := models.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	if err := r.db.Create(&newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to sign up %s", err)
	}
	return &newUser, nil
}

func (r *userRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, err
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	if err != nil {
		return nil, fmt.Errorf("no user found")
	}

	return users, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, fmt.Errorf("user not found by this email")
	}
	return &user, err
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, fmt.Errorf("user not found by this username")
	}
	return &user, err
}

func (r *userRepository) Update(id string, user *models.UpdateUser) (*models.User, error) {
	var existUser models.User

	if err := r.db.First(&existUser, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if err := r.db.Model(&existUser).Updates(user).Error; err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return &existUser, nil
}

func (r *userRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}

// func (r *userRepository) UpdatePassword(userID uuid.UUID, hashedPassword string) error {

// 	return err
// }

func (r *userRepository) UpdateRole(userID string, req models.UpdateRoleRequest) error {

	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"role":       req.Role,
			"updated_at": time.Now(),
		}).Error
}
