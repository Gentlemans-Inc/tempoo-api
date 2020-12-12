package main

import (
	"fmt"
	"log"
	"os"

	"github.com/api/router"
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

	router.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":5000"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(app.Listen(port))
}
