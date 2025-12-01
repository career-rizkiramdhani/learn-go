package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"crud/config"
	"crud/models"
)

// CreateUser creates a new user, hashing the password and generating a code
func CreateUser(username, password, email, prefix string, createdBy uint) (*models.User, error) {
	// check unique username/email
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", username).Or("email = ?", email).Count(&count)
	if count > 0 {
		return nil, errors.New("username or email already exists")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user := &models.User{
		Code:         GenerateUserCode(prefix),
		Username:     username,
		PasswordHash: string(pw),
		Email:        email,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    now,
		UpdatedAt:    now,
		Status:       "active",
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID returns a user by id
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates fields of an existing user. password may be empty to keep existing.
func UpdateUser(id uint, username, password, email, status *string, updatedBy uint) (*models.User, error) {
	user, err := GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if username != nil {
		user.Username = *username
	}
	if email != nil {
		user.Email = *email
	}
	if status != nil {
		user.Status = *status
	}
	if password != nil && *password != "" {
		pw, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(pw)
	}

	user.UpdatedBy = updatedBy
	user.UpdatedAt = time.Now()

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser soft-deletes a user by id
func DeleteUser(id uint) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ListUsers returns all users (simple implementation)
// ListUsers returns users with pagination. page starts from 1. size is number per page.
// It also returns total number of users.
func ListUsers(page, size int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	var users []models.User
	var total int64

	if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	if err := config.DB.Limit(size).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
