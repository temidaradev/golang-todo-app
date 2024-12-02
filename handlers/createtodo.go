package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/temidaradev/golang-todo-app/database"
	"github.com/temidaradev/golang-todo-app/models"
)

func CreateTodo(c *fiber.Ctx) error {
	fact := new(models.Todo)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Redirect("/")
}
