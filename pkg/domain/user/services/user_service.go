package services

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/model"
	userRepo "github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/repository"
)

// UserService interface
type UserService interface {
	CreateUser(user *model.User) (*model.Response, error)
	UpdateUser(user *model.User, id int) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (user model.User, err error)
	GetUserById(id int) (user model.User, err error)
}

// NewUserService returns a UserService implementation
func NewUserService(repository userRepo.Repository) (service UserService) {
	service = &Service{
		Repository: repository,
	}
	return
}
