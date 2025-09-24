package models

import (
	"time"
	"wallet-service/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `gorm:"unique" json:"email"`
	Role      string     `gorm:"default:'user'" json:"role"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate is a GORM hook that automatically hashes the user's password before creation.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = utils.HashPassword(u.Password)
	}
	return
}
