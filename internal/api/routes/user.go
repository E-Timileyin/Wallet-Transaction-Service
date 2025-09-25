package routes

import (
	"wallet-service/internal/api/handlers"
	"wallet-service/internal/api/middleware"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures the user-related routes.
func SetupUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler, authService *service.AuthService) {
	userRoutes := router.Group("/api/users")
	userRoutes.Use(middleware.AuthMiddleware(authService))
	{
		userRoutes.GET("", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("/email/:email", userHandler.GetUserByEmail)
		userRoutes.POST("", userHandler.CreateUser)
	}
}
