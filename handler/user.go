package handler

import (
	"github.com/api/models"
	"github.com/api/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser Handler for POST /user
func CreateUser(c *fiber.Ctx)  error {
	db := database.DB
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid fields"})
	}
	
	hash, err := hashPassword(user.Password)
	
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)	
	}
	
	user.Password = hash
	
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := models.NewUser{
		Email:    user.Email,
		Name: user.Name,
	}
	

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}