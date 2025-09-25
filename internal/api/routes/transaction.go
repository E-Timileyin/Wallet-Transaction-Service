package routes

import (
	"wallet-service/internal/api/handlers"
	"wallet-service/internal/api/middleware"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupTransactionRoutes configures the transaction-related routes.
func SetupTransactionRoutes(router *gin.Engine, transactionHandler *handlers.TransactionHandler, authService *service.AuthService) {
	transactionRoutes := router.Group("/api")
	transactionRoutes.Use(middleware.AuthMiddleware(authService))
	{
		transactionRoutes.GET("/wallets/:walletID/transactions", transactionHandler.GetTransactions)
	}
}
