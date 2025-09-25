package handlers

import (
	"net/http"
	"strconv"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
	_ "wallet-service/docs"
)

// TransactionHandler handles transaction-related HTTP requests.
type TransactionHandler struct {
	transactionService *service.TransactionService
}

// NewTransactionHandler creates a new TransactionHandler.
func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

// GetTransactions handles the request to get all transactions for a wallet.
// @Summary Get wallet transactions
// @Description Retrieve all transactions for a specific wallet
// @Tags Transactions
// @Security ApiKeyAuth
// @Produce json
// @Param walletID path int true "Wallet ID"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/wallets/{walletID}/transactions [get]
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	walletID, err := strconv.ParseUint(c.Param("walletID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	transactions, err := h.transactionService.GetTransactionsByWalletID(uint(walletID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
