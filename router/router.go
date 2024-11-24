package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/temidaradev/golang-todo-app/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", adaptor.HTTPHandler(handlers.Make(handlers.HandleHome)))
	app.Get("/db", handlers.ListFact)

	app.Post("/Facts", handlers.CreateFact)
}
