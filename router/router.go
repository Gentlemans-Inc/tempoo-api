package router

import (
	"github.com/api/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// Api base
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Health
	health := v1.Group("/health")
	health.Get("/", handler.HealthCheck)

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := v1.Group("/user")
	user.Post("/", handler.CreateUser)
}
