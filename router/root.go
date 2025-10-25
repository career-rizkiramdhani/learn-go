package router

import (
	"net/http"

	"crud/dto"
	"crud/service"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Group all routes under /api
	g := e.Group("/api")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// route untuk signin menggunakan service
	g.POST("/signin", signInHandler)
}

func signInHandler(c echo.Context) error {
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
