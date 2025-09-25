package models

import (
	"time"
	"wallet-service/pkg/utils"

	"gorm.io/gorm"
)

// User represents a user in the system
// @Description User model
type User struct {
	ID        uint       `json:"id" example:"1" gorm:"primaryKey"`
	Name      string     `json:"name" example:"John Doe"`
	Email     string     `json:"email" example:"john@example.com" gorm:"unique"`
	Role      string     `json:"role" example:"user" gorm:"default:'user'"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// BeforeCreate is a GORM hook that automatically hashes the user's password before creation.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = utils.HashPassword(u.Password)
	}
	return
}
