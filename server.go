package main

import (
	"github.com/api/database"
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
	// Database connection
	database.ConnectDatabase()

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

	//  I don't know why it's not recognizing .Close()
	// defer database.DB.Close()
}
