package main

import (
	"crud/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
