package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user account with audit fields
type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Code         string `gorm:"uniqueIndex;not null" json:"code"`
	Username     string `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string `gorm:"not null" json:"-"`
	Email        string `gorm:"uniqueIndex;not null" json:"email"`

	CreatedBy uint           `json:"created_by"`
	UpdatedBy uint           `json:"updated_by"`
	CreatedAt time.Time      `json:"created_time"`
	UpdatedAt time.Time      `json:"updated_time"`
	Status    string         `json:"status"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
