package router

import (
	"net/http"

	"crud/handler"
	"crud/middleware"

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

	// User CRUD
	g.POST("/users", handler.CreateUser)
	g.GET("/users", handler.ListUsers)
	g.GET("/users/:id", handler.GetUser)
	g.PUT("/users/:id", handler.UpdateUser)
	g.DELETE("/users/:id", handler.DeleteUser)

	// protected routes (require JWT)
	pg := g.Group("")
	pg.Use(middleware.RequireAuth)

	pg.POST("/logout", handler.Logout)
	pg.GET("/users", handler.ListUsers)
	pg.GET("/users/:id", handler.GetUser)
	pg.PUT("/users/:id", handler.UpdateUser)
	pg.DELETE("/users/:id", handler.DeleteUser)
}
