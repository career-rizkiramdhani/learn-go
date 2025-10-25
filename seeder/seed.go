package seeder

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"crud/config"
	"crud/models"
)

// Seed creates initial data (admin user) if none exists
func Seed() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("seeder: users already exist, skipping")
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	now := time.Now()

	user := models.User{
		Code:         "USR-1",
		Username:     "admin",
		PasswordHash: string(pw),
		Email:        "admin@example.com",
		CreatedBy:    0,
		UpdatedBy:    0,
		CreatedAt:    now,
		UpdatedAt:    now,
		Status:       "active",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Fatalf("seeder create user failed: %v", err)
	}
	log.Println("seeder: created admin user with password 'password'")
}
