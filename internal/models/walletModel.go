package models

// Wallet represents a user's wallet
// @Description Wallet model
type Wallet struct {
	ID        uint    `json:"id" example:"1"`
	CreatedAt string  `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt string  `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	UserID    uint    `json:"user_id" example:"1" gorm:"not null"`
	Name      string  `json:"name" example:"My Wallet" gorm:"not null"`
	Balance   float64 `json:"balance" example:"100.50" gorm:"default:0"`
	Currency  string  `json:"currency" example:"USD" gorm:"default:'USD'"`
	IsActive  bool    `json:"is_active" example:"true" gorm:"default:true"`
}

// FundWalletRequest defines the request body for funding a wallet.
// @Description Wallet funding request
type FundWalletRequest struct {
	Amount float64 `json:"amount" example:"100.50" binding:"required,gt=0"`
}
