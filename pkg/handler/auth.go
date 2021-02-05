package handler

import (
	"github.com/Mangaba-Labs/tempoo-api/internal/user"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Login Handler for POST /auth/login
func Login(c *fiber.Ctx) error {

	var service = user.NewUserService()
	var input user.AuthRequest
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(fiber.Map{"status": "error", "error": "malformed auth request", "data": nil})

	}
	email := input.Email
	pass := input.Password

	usr, err := service.GetUserByEmail(email)

	if err != nil || len(usr.Email) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pass)); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
