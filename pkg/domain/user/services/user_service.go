package services

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/database"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/repository"
)

type UserService interface {
	CreateUser(user *user.User) (*user.Response, error)
	UpdateUser(user *user.User, id int) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (user user.User, err error)
	GetUserById(id int) (user user.User, err error)
}

// Should have a comment
func NewUserService() (service UserService) {
	r := repository.NewUserRepository(database.Instance)
	service = &Service{
		Repository: r,
	}
	return
}
