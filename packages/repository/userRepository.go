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
	GetByID(id uuid.UUID) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(id uuid.UUID, user *models.UpdateUser) (*models.User, error)
	Delete(id uuid.UUID) error
	// UpdatePassword(userID uuid.UUID, hashedPassword string) error
	UpdateRole(userID uuid.UUID, role models.Role) error
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
		ID:       uuid.New(),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}
	if err := r.db.Create(&newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to sign up %s", err)
	}
	return &newUser, nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, err
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

func (r *userRepository) Update(id uuid.UUID, user *models.UpdateUser) (*models.User, error) {
	var existUser models.User

	if err := r.db.First(&existUser, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if err := r.db.Model(&existUser).Updates(user).Error; err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return &existUser, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, id).Error
}

// func (r *userRepository) UpdatePassword(userID uuid.UUID, hashedPassword string) error {

// 	return err
// }

func (r *userRepository) UpdateRole(userID uuid.UUID, role models.Role) error {

	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"role":       role,
			"updated_at": time.Now(),
		}).Error
}
