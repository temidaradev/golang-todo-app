package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/temidaradev/golang-todo-app/database"
	"github.com/temidaradev/golang-todo-app/models"
	"github.com/temidaradev/golang-todo-app/views/home"
)

func HandleHome(c *fiber.Ctx) error {
	dbTodo := []models.Todo{}
	viewTodo := []home.Task{}

	result := database.DB.Db.Find(&dbTodo)
	if result.Error != nil {
		return result.Error
	}

	for _, t := range dbTodo {
		homeTodo := home.Task{
			ID:   fmt.Sprint(t.ID),
			Todo: t.Todo,
		}
		viewTodo = append(viewTodo, homeTodo)
	}

	return Render(c, home.Index(viewTodo))
}
