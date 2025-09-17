package main

import (
	"log"
	"wallet-service/internal/api"
	"wallet-service/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	gin.SetMode(gin.DebugMode)
	router := api.NewRouter()

	log.Println("🚀 Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
