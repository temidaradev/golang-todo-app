package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome!")
	})

	listenAddr := os.Getenv("LISTEN_ADDR")

	log.Fatal(app.Listen(listenAddr))
}
