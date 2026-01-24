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
	g.POST("/logout", handler.Logout)

	// User CRUD
	g.POST("/users", handler.CreateUser)
	g.GET("/users", handler.ListUsers)
	g.GET("/users/:id", handler.GetUser)
	g.PUT("/users/:id", handler.UpdateUser)
	g.DELETE("/users/:id", handler.DeleteUser)
}
