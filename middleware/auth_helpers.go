package middleware

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// GetUserID retrieves `user_id` from the Echo context and returns it as uint64.
// The second return value indicates whether a valid id was present.
func GetUserID(c echo.Context) (uint64, bool) {
	v := c.Get("user_id")
	if v == nil {
		return 0, false
	}
	switch val := v.(type) {
	case uint64:
		return val, true
	case uint:
		return uint64(val), true
	case int:
		return uint64(val), true
	case int64:
		return uint64(val), true
	case float64:
		return uint64(val), true
	case string:
		if id, err := strconv.ParseUint(val, 10, 64); err == nil {
			return id, true
		}
	}
	return 0, false
}

// GetUserClaims returns JWT claims stored in context by the auth middleware.
func GetUserClaims(c echo.Context) (jwt.MapClaims, bool) {
	v := c.Get("user_claims")
	if v == nil {
		return nil, false
	}
	if claims, ok := v.(jwt.MapClaims); ok {
		return claims, true
	}
	return nil, false
}
