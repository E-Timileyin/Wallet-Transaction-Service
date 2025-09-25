package repository

import (
	"wallet-service/internal/db"
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

// WalletRepository handles database operations for wallets.
type WalletRepository struct {
	DB *gorm.DB
}

// NewWalletRepository creates a new WalletRepository.
func NewWalletRepository() *WalletRepository {
	return &WalletRepository{DB: db.DB}
}

// Create creates a new wallet in the database.
func (r *WalletRepository) Create(wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}

// FindByUserID finds a wallet by the user's ID.
func (r *WalletRepository) FindByUserID(userID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Update updates a wallet's details in the database.
func (r *WalletRepository) Update(wallet *models.Wallet) error {
	return r.DB.Save(wallet).Error
}
