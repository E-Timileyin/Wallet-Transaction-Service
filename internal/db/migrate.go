package db

import (
	"wallet-service/internal/db/migrations"
)

func RunMigrations() {
	migrations.MigrateUsers(DB)
	migrations.MigrateWallets(DB)
	migrations.MigrateTransactions(DB)
}
