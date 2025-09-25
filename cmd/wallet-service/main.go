package main

import (
	"log"
	"wallet-service/internal/api/handlers"
	"wallet-service/internal/api/routes"
	"wallet-service/internal/config"
	"wallet-service/internal/db"
	"wallet-service/internal/repository"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "wallet-service/docs"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v", err)
	}

	// Initialize database
	db.InitDB()

	// Set up Gin mode
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// Initialize components
	userRepo := repository.NewUserRepository()
	walletRepo := repository.NewWalletRepository()
	transactionRepo := repository.NewTransactionRepository()

	walletService := service.NewWalletService(walletRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	authService := service.NewAuthService(*userRepo, walletService, &cfg)
	userService := service.NewUserService(userRepo)
	adminService := service.NewAdminService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	adminHandler := handlers.NewAdminHandler(adminService)
	walletHandler := handlers.NewWalletHandler(walletService, transactionService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Setup routes
	routes.SetupAuthRoutes(router, authHandler)
	routes.SetupUserRoutes(router, userHandler, authService)
	routes.SetupAdminRoutes(router, adminHandler, authService, userRepo)
	routes.SetupWalletRoutes(router, walletHandler, authService)
	routes.SetupTransactionRoutes(router, transactionHandler, authService)

	// Swagger UI route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	// Start server
	log.Println("🚀 Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
