package main

import (
	"boilerplate-echo-gorm-postgres/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo

	e := echo.New()

	routes.InitRoutes(e)
	// Middleware
	e.Use(middleware.Logger())

	e.Start(":8080")

}
