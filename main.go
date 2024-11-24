package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/temidaradev/golang-todo-app/database"
	"github.com/temidaradev/golang-todo-app/router"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router.SetupRoutes(app)

	listenAddr := os.Getenv("LISTEN_ADDR")

	log.Fatal(app.Listen(listenAddr))
}
