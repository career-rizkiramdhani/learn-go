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
