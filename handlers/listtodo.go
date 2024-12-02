package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/temidaradev/golang-todo-app/database"
	"github.com/temidaradev/golang-todo-app/models"
)

func ListTodo(c *fiber.Ctx) error {
	facts := []models.Todo{}

	result := database.DB.Db.Find(&facts)
	if result.Error != nil {
		return result.Error
	}

	return c.Status(200).JSON(facts)
}
