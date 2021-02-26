package handler

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/services"
	"github.com/gofiber/fiber/v2"
)

// UserHandler interface
type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	EditUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

// NewUserHandler returns a pointer to an handler impl
func NewUserHandler(s services.UserService) Handler {
	return Handler{
		service: s,
	}
}
