package models

import "time"

type User struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	IsAdmin        bool      `json:"is_admin" gorm:"default:false"`
	ProfilePicture *string   `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
