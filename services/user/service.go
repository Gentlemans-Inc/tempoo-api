package user

import (
	"github.com/api/models"
	"github.com/api/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository repository.UserRepository
}

// CreateUser on app
func (s Service) CreateUser(user *models.User) (*models.NewUser, error) {
	_, err := s.Repository.FindOneByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	hash, err := hashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hash

	if err := s.Repository.Create(user); err != nil {
		return nil, err
	}

	return &models.NewUser{
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s Service) UpdateUser(user models.User) error {
	panic("implement me")
}

func (s Service) DeleteUser(id int) error {
	panic("implement me")
}

func (s Service) GetUserByEmail(email string) (user models.User, err error) {
	panic("implement me")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
