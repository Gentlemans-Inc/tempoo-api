package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	//Handle Cors
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Printf("Listening on port %s\n\n", port)
	app.Listen(port)
}
