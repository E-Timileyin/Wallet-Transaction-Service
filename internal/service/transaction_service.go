package service

import (
	"wallet-service/internal/models"
	"wallet-service/internal/repository"
)

// TransactionService provides transaction-related services.
type TransactionService struct {
	transactionRepo *repository.TransactionRepository
}

// NewTransactionService creates a new TransactionService.
func NewTransactionService(transactionRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

// GetTransactionsByWalletID retrieves all transactions for a specific wallet.
func (s *TransactionService) GetTransactionsByWalletID(walletID uint) ([]models.Transaction, error) {
	return s.transactionRepo.FindByWalletID(walletID)
}

// CreateTransaction creates a new transaction record.
func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	return s.transactionRepo.Create(transaction)
}
