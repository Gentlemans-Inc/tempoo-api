package handler

import (
	"github.com/api/internal/weather"
	"github.com/gofiber/fiber/v2"
)

// WeatherHandler handler for /weather endpoints
type WeatherHandler interface {
	GetWeather(c *fiber.Ctx) error
}

// ServiceHandler struct
type ServiceHandler struct {
	service weather.WeatherService
}

// NewWeatherHandler initializer
func NewWeatherHandler(s weather.WeatherService) WeatherHandler {
	return &ServiceHandler{
		service: s,
	}
}

// GetWeather handler for GET /weather/current
func (h *ServiceHandler) GetWeather(c *fiber.Ctx) error {
	var userCoord = new(weather.Request)
	var service = weather.NewWeatherService()

	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		return c.JSON(fiber.Map{"status": "error", "error": "malformed get-weather request", "data": nil})
	}

	userCoord.Latitude = lat
	userCoord.Longitude = lon

	currentWeather, err := service.GetCurrentWeather(userCoord)

	if err != nil {
		if err.Error() == "400" {
			c.Context().Response.SetStatusCode(400)
			return c.JSON(fiber.Map{"status": "error", "error": "malformed get-weather request", "data": nil})
		} else if err.Error() == "500" {
			c.Context().Response.SetStatusCode(500)
			return c.JSON(fiber.Map{"status": "error", "data": nil})
		} else {
			return c.JSON(fiber.Map{"status": "error", "data": nil})
		}
	}

	return c.JSON(fiber.Map{"status": "success", "data": currentWeather})
}
