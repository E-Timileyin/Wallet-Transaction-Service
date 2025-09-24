package routes

import (
	"wallet-service/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configures the authentication routes.
func SetupAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.RefreshToken)
	}
}
