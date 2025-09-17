// 0003_create_transactions.go
package migrations

import (
	"wallet-service/internal/models"

	"gorm.io/gorm"

	"log"
)

func MigrateTransactions(db *gorm.DB) {
	err := db.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatalf("Failed to migrate transactions table: %v", err)
	}
	log.Println("✅ Transactions table migrated")
}
