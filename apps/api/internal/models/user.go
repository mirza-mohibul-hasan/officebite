package models

import "time"

type UserRole string

const (
	RoleEmployee UserRole = "employee"
	RoleAdmin    UserRole = "admin"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:120;not null" json:"name"`
	Email        string    `gorm:"size:160;not null;uniqueIndex" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         UserRole  `gorm:"type:varchar(24);not null;default:'employee'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
