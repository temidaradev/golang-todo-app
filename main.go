package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/joho/godotenv"
	"github.com/temidaradev/golang-todo-app/database"
	"github.com/temidaradev/golang-todo-app/router"
)

func main() {
	database.ConnectDB()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	listenAddr := os.Getenv("LISTEN_ADDR")

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
	router.SetupRoutes(app)
	app.Static("/*", "./public/styles.css")

	go func() {
		if err := app.Listen(listenAddr); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c //blocking the main thread until interrupted

	app.Shutdown()
	fmt.Println("shutting down the server")
}
