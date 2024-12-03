package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/temidaradev/golang-todo-app/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.HandleHome)
	app.Get("/db", handlers.ListTodo)
	app.Post("/Delete/:id", handlers.DeleteTodo)

	app.Post("/Todos", handlers.CreateTodo)
}
