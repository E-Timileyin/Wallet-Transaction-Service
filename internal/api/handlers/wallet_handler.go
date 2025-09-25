package handlers

import (
	"log"
	"net/http"
	"wallet-service/internal/models"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
	_ "wallet-service/docs"
)

// WalletHandler handles wallet-related HTTP requests.
type WalletHandler struct {
	walletService      *service.WalletService
	transactionService *service.TransactionService
}

// NewWalletHandler creates a new WalletHandler.
func NewWalletHandler(walletService *service.WalletService, transactionService *service.TransactionService) *WalletHandler {
	return &WalletHandler{walletService: walletService, transactionService: transactionService}
}

// GetWallet handles the request to get the current user's wallet.
// @Summary Get user wallet
// @Description Retrieve the current user's wallet information
// @Tags Wallets
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.Wallet
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/wallet [get]
func (h *WalletHandler) GetWallet(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	wallet, err := h.walletService.GetWalletByUserID(uint(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// FundWallet handles the request to add funds to the user's wallet.
// @Summary Fund wallet
// @Description Add funds to the user's wallet
// @Tags Wallets
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body models.FundWalletRequest true "Funding request"
// @Success 200 {object} models.Wallet
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/wallet/fund [post]
func (h *WalletHandler) FundWallet(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	var req models.FundWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the wallet to find its ID and current balance
	wallet, err := h.walletService.GetWalletByUserID(uint(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	// Fund the wallet
	updatedWallet, err := h.walletService.FundWallet(uint(userID.(float64)), req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fund wallet"})
		return
	}

	// Create a transaction record
	transaction := &models.Transaction{
		WalletID:       wallet.ID,
		Type:           "deposit",
		Amount:         req.Amount,
		Description:    "Wallet funding",
		Status:         "completed",
		Reference:      "",
		BalanceBefore:  wallet.Balance,
		BalanceAfter:   updatedWallet.Balance,
	}

	if err := h.transactionService.CreateTransaction(transaction); err != nil {
		// Log the error, but don't fail the funding
		log.Printf("Failed to create transaction for wallet %d: %v", wallet.ID, err)
	}

	c.JSON(http.StatusOK, updatedWallet)
}
