package handler

import (
	"net/http"

	"crud/dto"
	"crud/service"

	"github.com/labstack/echo/v4"
)

// SignIn handles the user login request and returns token + user info
func SignIn(c echo.Context) error {
	var req dto.SignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, token, err := service.Authenticate(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	resp := dto.SignInFullResponse{
		Code:    200,
		Message: "Success",
		Data: dto.SignInData{
			Result: dto.SignInResult{
				User: dto.SignInUser{
					ID:       user.ID,
					Username: user.Username,
					Email:    user.Email,
				},
				Token: token,
			},
		},
	}

	return c.JSON(http.StatusOK, resp)
}
