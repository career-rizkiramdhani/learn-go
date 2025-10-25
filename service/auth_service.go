package service

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"crud/config"
	"crud/models"
)

// Authenticate checks username/password and returns a signed JWT when valid
func Authenticate(username, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.Cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
