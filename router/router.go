package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/temidaradev/golang-todo-app/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.HandleHome)
	app.Get("/db", handlers.ListFact)

	app.Post("/Facts", handlers.CreateFact)
}
