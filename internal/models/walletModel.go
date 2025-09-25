package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserID    uint    `gorm:"not null" json:"user_id"`
	Name      string  `gorm:"not null" json:"name"`
	Balance   float64 `gorm:"default:0" json:"balance"`
	Currency  string  `gorm:"default:'USD'" json:"currency"`
	IsActive  bool    `gorm:"default:true" json:"is_active"`
}

// FundWalletRequest defines the request body for funding a wallet.
type FundWalletRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
