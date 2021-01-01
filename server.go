package main

import (
	"fmt"
	"log"
	"os"

	"github.com/api/database"

	"github.com/api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Database connection
	database.ConnectDatabase()

	app := fiber.New()

	//Helmet security
	//app.Use(helmet.New())

	//Handle Cors
	app.Use(cors.New())

	//Rate limiting
	app.Use(limiter.New())

	//Handle panics
	app.Use(recover.New())

	//Handle logs
	app.Use(logger.New())

	//Request ID
	//app.Use(requestid.New())

	//Handle routes
	router.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(app.Listen(port))

}
