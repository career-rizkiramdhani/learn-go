package auth

import "time"

// AuthToken represents an issued authentication token
type AuthToken struct {
	Token     string
	ExpiresAt time.Time
}
