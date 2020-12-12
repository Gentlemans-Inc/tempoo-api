package handler

import (
	"time"

	"github.com/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Login Handler for POST /auth/login
func Login(c *fiber.Ctx) error {

	var input models.UserLogin
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)

	}
	email := input.Email
	pass := input.Password
	if email != "arthur" || pass != "arthur" {
		return c.SendStatus(fiber.StatusUnauthorized)

	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)

	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
