package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"wallet-service/internal/models"
)

var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB() {
	// Load .env first (if exists)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, relying on system env variables")
	}

	// Read DATABASE_URL from env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL env variable not set")
	}

	// Connect to Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")

	// Auto-migrate your models
	if err := DB.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{}); err != nil {
		log.Fatalf("❌ Failed to migrate tables: %v", err)
	}
	log.Println("✅ Database tables migrated successfully")
}
