// 0001_create_users.go
package migrations

import (
	"gorm.io/gorm"
	"wallet-service/internal/models"

	"log"
)

func MigrateUsers(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate users table: %v", err)
	}
	log.Println("✅ Users table migrated")
}
