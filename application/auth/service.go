package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"crud/config"
	userdomain "crud/domain/user"
	"crud/infrastructure/persistence"
)

// SignIn validates credentials and returns JWT token string and expiry
func SignIn(username, password string) (string, time.Time, *userdomain.User, error) {
	// use repository to find user by username
	repo := persistence.NewGormUserRepository()
	u, err := repo.FindByUsername(username)
	if err != nil {
		return "", time.Time{}, nil, err
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", time.Time{}, nil, err
	}

	// create JWT token
	expiresAt := time.Now().Add(time.Duration(config.Cfg.JWTExpiredHours) * time.Hour)
	claims := jwt.MapClaims{
		"sub": u.ID,
		"exp": expiresAt.Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.Cfg.JWTSecret))
	if err != nil {
		return "", time.Time{}, nil, err
	}

	return signed, expiresAt, u, nil
}

// Logout is a placeholder; for stateless JWT, logout is client-side (or add blacklist)
func Logout(tokenString string) error {
	// optional: implement token blacklist in infrastructure if needed
	return nil
}
