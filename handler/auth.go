package handler

import (
	"net/http"

	"crud/dto"
	"crud/service"

	"github.com/labstack/echo/v4"
)

// SignIn handles the user login request
func SignIn(c echo.Context) error {
	var req dto.SignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, err := service.Authenticate(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, dto.SignInResponse{Token: token})
}
