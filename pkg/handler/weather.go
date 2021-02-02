package handler

import "github.com/gofiber/fiber/v2"

func GetWeather(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "data": "27 graus"})
}