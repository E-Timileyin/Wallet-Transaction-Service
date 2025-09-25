package models

// Transaction represents a wallet transaction
// @Description Transaction model for wallet operations
type Transaction struct {
	ID            uint    `json:"id" example:"1"`
	CreatedAt     string  `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt     string  `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt     *string `json:"deleted_at,omitempty"`
	WalletID      uint    `json:"wallet_id" example:"1" gorm:"not null"`
	Type          string  `json:"type" example:"deposit" gorm:"not null"` // "deposit", "withdrawal", "transfer"
	Amount        float64 `json:"amount" example:"100.50" gorm:"not null"`
	Description   string  `json:"description" example:"Wallet funding"`
	Status        string  `json:"status" example:"completed" gorm:"default:pending"` // "pending", "completed", "failed"
	Reference     string  `json:"reference" example:"REF123456" gorm:"unique"`
	BalanceBefore float64 `json:"balance_before" example:"0.00"`
	BalanceAfter  float64 `json:"balance_after" example:"100.50"`
}
