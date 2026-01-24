package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"crud/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// RequireAuth validates JWT from Authorization header, checks expiry, and
// places `user_id` and `user_claims` into the Echo context for handlers.
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing authorization"})
		}

		if !strings.HasPrefix(auth, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header"})
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
		}

		// check expiration if present
		if expv, ok := claims["exp"]; ok {
			var exp int64
			switch v := expv.(type) {
			case float64:
				exp = int64(v)
			case int64:
				exp = v
			case string:
				if p, err := strconv.ParseInt(v, 10, 64); err == nil {
					exp = p
				}
			}
			if exp != 0 && time.Unix(exp, 0).Before(time.Now()) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "token expired"})
			}
		}

		// extract subject (user id)
		var userID uint64
		if sub, ok := claims["sub"]; ok {
			switch v := sub.(type) {
			case float64:
				userID = uint64(v)
			case string:
				if id, err := strconv.ParseUint(v, 10, 64); err == nil {
					userID = id
				}
			}
		}
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token subject"})
		}

		c.Set("user_id", userID)
		c.Set("user_claims", claims)
		return next(c)
	}
}
