package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleRecruiter Role = "recruiter"
	RoleAdmin     Role = "admin"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username       string    `json:"username" validate:"required,min=3,max=10"`
	Password       string    `json:"password" validate:"required,min=4,max=10"`
	Email          string    `json:"email" validate:"required"`
	Jobs           []Job     `json:"jobs" gorm:"foreignKey:UserID"`
	Role           Role      `json:"role" validate:"required,oneof=user recruiter admin"`
	ProfilePicture *string   `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type JobType string

const (
	Remote JobType = "remote"
	Hybrid JobType = "hybrid"
	Onsite JobType = "onsite"
)

type Job struct {
	Id              uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title           string    `json:"title" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	Location        string    `json:"location" validate:"required"`
	Company         string    `json:"company" validate:"required"`
	MinSalary       *int      `json:"min_salary" validate:"required"`
	ExperienceLevel string    `json:"experience_level" validate:"required"`
	Skills          string    `json:"skills" validate:"required"`
	MaxSalary       *int      `json:"max_salary" validate:"required"`
	Type            JobType   `json:"type" validate:"required,oneof=remote hybrid onsite"`
	CreatedAt       time.Time `json:"created_at"`
	ApplyLink       *string   `json:"applyLink"`

	UserID uuid.UUID `json:"userId" validate:"required"` // foreign key
	User   User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateUser struct {
	Username       string  `json:"username" validate:"required,min=3,max=10"`
	Password       string  `json:"password" validate:"required,min=4,max=10"`
	Email          string  `json:"email" validate:"required"`
	Role           Role    `json:"role" validate:"required,oneof=user recruiter admin"`
	ProfilePicture *string `json:"profile_picture"`
}

type UpdateUser struct {
	Username       *string `json:"username" validate:"omitempty,min=3,max=10"`
	Email          *string `json:"email"`
	ProfilePicture *string `json:"profile_picture"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required,min=3,max=10"`
	Password string `json:"password" validate:"required,min=4,max=10"`
	Email    string `json:"email" validate:"required"`
}

type JWTClaims struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
}
