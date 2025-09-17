package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	WalletID      uint    `gorm:"not null"`
	Type          string  `gorm:"not null"` // "deposit", "withdrawal", "transfer"
	Amount        float64 `gorm:"not null"`
	Description   string
	Status        string `gorm:"default:pending"` // "pending", "completed", "failed"
	Reference     string `gorm:"unique"`
	BalanceBefore float64
	BalanceAfter  float64
}
