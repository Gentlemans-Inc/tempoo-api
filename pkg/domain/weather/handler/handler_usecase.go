package handler

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/model"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/services"
	"github.com/gofiber/fiber/v2"
)

// ServiceHandler struct
type ServiceHandler struct {
	service services.WeatherService
}

// GetWeather handler for GET /weather/current
func (h *ServiceHandler) GetWeather(c *fiber.Ctx) error {
	var userCoord = &model.Request{}
	var service = services.NewWeatherService()

	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		return c.JSON(fiber.Map{"status": "error", "message": "malformed get-weather request", "data": nil})
	}

	userCoord.Latitude = lat
	userCoord.Longitude = lon

	currentWeather, err := service.GetCurrentWeather(userCoord)

	if err != nil {
		if err.Error() == "400" {
			c.Context().Response.SetStatusCode(400)
			return c.JSON(fiber.Map{"status": "error", "message": "malformed get-weather request", "data": nil})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": nil, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": currentWeather})
}
