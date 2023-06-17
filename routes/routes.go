package routes

import (
	"boilerplate-echo-gorm-postgres/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/tasks", controllers.TaskPost)
	e.GET("/tasks", controllers.TaskGet)
	e.PUT("/tasks/:id", controllers.TasksPut)
	e.DELETE("/tasks/:id", controllers.TasksDel)
}
