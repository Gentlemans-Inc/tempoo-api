package services

import (
	"github.com/api/models"
	"github.com/api/repository"
	"github.com/api/services/user"
)

type UserService interface {
	CreateUser(user *models.User) (*models.NewUser, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (user models.User, err error)
}


func NewUserService() (service UserService) {
	r := repository.NewUserRepository()
	service = &user.Service{
		Repository: r,
	}
	return
}