package routes

import (
	"wallet-service/internal/api/handlers"
	"wallet-service/internal/api/middleware"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupWalletRoutes configures the wallet-related routes.
func SetupWalletRoutes(router *gin.Engine, walletHandler *handlers.WalletHandler, authService *service.AuthService) {
	walletRoutes := router.Group("/api/wallet")
	walletRoutes.Use(middleware.AuthMiddleware(authService))
	{
		walletRoutes.GET("", walletHandler.GetWallet)
		walletRoutes.POST("/fund", walletHandler.FundWallet)
	}
}
