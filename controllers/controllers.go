package controllers

import (
	"boilerplate-echo-gorm-postgres/db/connection"
	"boilerplate-echo-gorm-postgres/db/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db = connection.ConnectToDB()

func TaskGet(c echo.Context) error {
	var tasks []models.Task
	db.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

// POST
func TaskPost(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return err
	}
	db.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

// PUT
func TasksPut(c echo.Context) error {
	id := c.Param("id")
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Task not found in PUT")
	}
	if err := c.Bind(&task); err != nil {
		return err
	}
	db.Save(&task)
	return c.JSON(http.StatusOK, task)
}

// DELETE
func TasksDel(c echo.Context) error {
	id := c.Param("id")
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Task not found in DELETE")
	}
	db.Delete(&task)
	return c.NoContent(http.StatusNoContent)
}
