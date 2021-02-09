package router

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/api/handler"
	middleware "github.com/Mangaba-Labs/tempoo-api/pkg/api/middlewares"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/services"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router pkg
func SetupRoutes(app *fiber.App) {

	userService := services.NewUserService()
	userHandler := handler.NewUserHandler(userService)
	weatherService := weather.NewWeatherService()
	weatherHandler := handler.NewWeatherHandler(weatherService)

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
	user := v1.Group("/users")
	user.Post("/", userHandler.CreateUser)
	user.Get("/:id", userHandler.GetUser)
	user.Delete("/:id", userHandler.DeleteUser)
	user.Put("/:id", userHandler.EditUser)

	// weather
	weather := v1.Group("/weather")
	weather.Get("/current", middleware.Protected(), weatherHandler.GetWeather)
}
