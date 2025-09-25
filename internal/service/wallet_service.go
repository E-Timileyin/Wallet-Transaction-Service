package service

import (
	"wallet-service/internal/models"
	"wallet-service/internal/repository"
)

// WalletService provides wallet-related services.
type WalletService struct {
	walletRepo *repository.WalletRepository
}

// NewWalletService creates a new WalletService.
func NewWalletService(walletRepo *repository.WalletRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo}
}

// CreateWallet creates a new wallet for a user.
func (s *WalletService) CreateWallet(userID uint, name string) (*models.Wallet, error) {
	wallet := &models.Wallet{
		UserID: userID,
		Name:   name,
	}
	if err := s.walletRepo.Create(wallet); err != nil {
		return nil, err
	}
	return wallet, nil
}

// GetWalletByUserID retrieves a wallet for a specific user.
func (s *WalletService) GetWalletByUserID(userID uint) (*models.Wallet, error) {
	return s.walletRepo.FindByUserID(userID)
}

// FundWallet adds funds to a user's wallet.
func (s *WalletService) FundWallet(userID uint, amount float64) (*models.Wallet, error) {
	wallet, err := s.walletRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	wallet.Balance += amount
	if err := s.walletRepo.Update(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}
