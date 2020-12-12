package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	//Handle Cors
	app.Use(cors.New())

	//Rate limiting
	app.Use(limiter.New())

	//Handle panics
	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = ":5000"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(app.Listen(port))
}
