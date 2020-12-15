package handler

import (
	"fmt"

	"github.com/api/database"
	"github.com/api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser Handler for POST /user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid fields"})
	}

	userExists := db.Find(&user, user.Email)
	if userExists != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User already exists"})
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
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

//EditUser handler for PUT /user/:id
func EditUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	fmt.Println()

	var newUser models.NewUser
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Invalid fields"})
	}

	var user models.User

	db.First(&user, id)
	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	user.Email = newUser.Email
	user.Name = newUser.Name

	result := db.Save(&user)
	if result == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error on saving user in database"})
	}

	editedUser := models.NewUser{
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "UpdatedUser", "data": editedUser})
}

// GetUser Handler for GET /user/:id
func GetUser(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB

	var user models.User

	db.Find(&user, id)

	userDetails := models.UserDetails{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	return c.JSON(fiber.Map{"status": "success", "user": userDetails})
}

// DeleteUser delete user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var user models.User

	db.First(&user, id)

	if user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	db.Delete(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted"})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}