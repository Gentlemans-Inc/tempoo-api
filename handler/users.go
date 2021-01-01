package handler

import (
	"github.com/api/database"
	"github.com/api/models"
	"github.com/api/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)




// CreateUser Handler for POST /user
func CreateUser(c *fiber.Ctx) error {
	var service = services.NewUserService()
	var user = new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": err})
	}

	newUser, err := service.CreateUser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

//EditUser handler for PUT /user/:id
func EditUser(c *fiber.Ctx) error {
	db := database.Instance
	id := c.Params("id")

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
	db := database.Instance
	id := c.Params("id")

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
	db := database.Instance
	id := c.Params("id")

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
