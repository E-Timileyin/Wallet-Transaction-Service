// 0002_create_wallets.go
package migrations

import (
	"wallet-service/internal/models"

	"gorm.io/gorm"

	"log"
)

func MigrateWallets(db *gorm.DB) {
	err := db.AutoMigrate(&models.Wallet{})
	if err != nil {
		log.Fatalf("Failed to migrate wallets table: %v", err)
	}
	log.Println("✅ Wallets table migrated")
}
