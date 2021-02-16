package router

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/api/handler"
	middleware "github.com/Mangaba-Labs/tempoo-api/pkg/api/middlewares"
	userHandler "github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/handler"
	weatherHandler "github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/handler"
	"github.com/gofiber/fiber/v2"
)


type Server struct {
	userHandler userHandler.Handler
	weatherHandler weatherHandler.WeatherHandler
}


func NewServer(userHandler userHandler.Handler, weatherHandler weatherHandler.WeatherHandler) *Server {
	return &Server{userHandler: userHandler, weatherHandler: weatherHandler}
}

// SetupRoutes setup router pkg
func (s *Server) SetupRoutes(app *fiber.App) {

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
	user.Post("/", s.userHandler.CreateUser)
	user.Get("/:id", s.userHandler.GetUser)
	user.Delete("/:id", s.userHandler.DeleteUser)
	user.Put("/:id", s.userHandler.EditUser)

	// weather
	weather := v1.Group("/weather")
	weather.Get("/current", middleware.Protected(), s.weatherHandler.GetWeather)
}
