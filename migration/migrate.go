package migration

import (
	"log"

	"crud/config"
	"crud/models"
)

// RunMigrations runs database migrations (GORM AutoMigrate)
func RunMigrations() {
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migrations failed: %v", err)
	}
}

// RollbackAll drops all known tables (use with caution)
func RollbackAll() {
	// Drop tables in reverse order of dependencies if needed
	if err := config.DB.Migrator().DropTable(&models.User{}); err != nil {
		log.Fatalf("rollback failed: %v", err)
	}
	log.Println("rollback: all tables dropped")
}
