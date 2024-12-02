package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/temidaradev/golang-todo-app/database"
	"github.com/temidaradev/golang-todo-app/models"
)

func DeleteTodo(c *fiber.Ctx) error {
	dbTodo := models.Todo{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&dbTodo)
	if result.Error != nil {
		return result.Error
	}

	return c.Redirect("/")
}
