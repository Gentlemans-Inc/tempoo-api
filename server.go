package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World")
	})

	app.Listen(":3000")
}
