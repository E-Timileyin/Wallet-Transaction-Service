package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
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
}
