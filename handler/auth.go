package handler

import (
	"net/http"

	authapp "crud/application/auth"
	"crud/dto"

	"github.com/labstack/echo/v4"
)

// SignIn handler
func SignIn(c echo.Context) error {
	var req dto.SignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, _, user, err := authapp.SignIn(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	res := dto.SignInResult{
		User: dto.SignInUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}

	return c.JSON(http.StatusOK, res)
}

// Logout handler (stateless token; returns success)
func Logout(c echo.Context) error {
	// optional: read token and revoke if blacklist implemented
	return c.JSON(http.StatusOK, dto.SignInResponse{Token: ""})
}
