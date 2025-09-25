package routes

import (
	"wallet-service/internal/api/handlers"
	"wallet-service/internal/api/middleware"
	"wallet-service/internal/repository"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupAdminRoutes configures the admin-only routes.
func SetupAdminRoutes(router *gin.Engine, adminHandler *handlers.AdminHandler, authService *service.AuthService, userRepo *repository.UserRepository) {
	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use(middleware.AuthMiddleware(authService))
	adminRoutes.Use(middleware.AdminMiddleware(userRepo))
	{
		adminRoutes.GET("/users", adminHandler.GetAllUsers)
		adminRoutes.PUT("/users/:id/role", adminHandler.UpdateUserRole)
		adminRoutes.DELETE("/users/:id", adminHandler.DeleteUser)
		adminRoutes.DELETE("/users", adminHandler.DeleteAllUsers)
	}
}
