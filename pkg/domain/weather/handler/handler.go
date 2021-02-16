package handler

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/services"
	"github.com/gofiber/fiber/v2"
)

type WeatherHandler interface {
	GetWeather(c *fiber.Ctx) error
}

// NewWeatherHandler initializer
func NewWeatherHandler(s services.WeatherService) WeatherHandler {
	return &ServiceHandler{
		service: s,
	}
}