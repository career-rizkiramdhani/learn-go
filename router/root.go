package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// baru: route untuk signin
	e.POST("/signin", signInHandler)
}

type signInRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func signInHandler(c echo.Context) error {
	var req signInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// contoh validasi sederhana
	if req.Username == "admin" && req.Password == "password" {
		return c.JSON(http.StatusOK, map[string]string{"token": "dummy-token"})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
}
