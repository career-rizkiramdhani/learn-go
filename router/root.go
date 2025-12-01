package router

import (
	"net/http"

	"crud/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	// Group all routes under /api
	g := e.Group("/api")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// route untuk signin menggunakan service
	g.POST("/signin", handler.SignIn)
}
