package repository

import (
	"wallet-service/internal/db"
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

// TransactionRepository handles database operations for transactions.
type TransactionRepository struct {
	DB *gorm.DB
}

// NewTransactionRepository creates a new TransactionRepository.
func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{DB: db.DB}
}

// Create creates a new transaction in the database.
func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

// FindByWalletID finds all transactions for a given wallet ID.
func (r *TransactionRepository) FindByWalletID(walletID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Where("wallet_id = ?", walletID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
