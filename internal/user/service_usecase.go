package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository UserRepository
}

// CreateUser on app
func (s Service) CreateUser(user *User) (*Response, error) {
	_, err := s.Repository.FindOneByEmail(user.Email)

	// err == nil means that we find an user with this e-mail on db
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hash, err := hashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hash

	if err := s.Repository.Create(user); err != nil {
		return nil, err
	}

	return &Response{
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s Service) UpdateUser(user *User, id int) (err error) {
	result, err := s.Repository.FindOneByEmail(user.Email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %s on database", user.Email)
		return errors.New(errMessage)
	}

	result.Name = user.Name
	result.Email = user.Email

	err = s.Repository.Create(&result)
	return
}

func (s Service) UpdateUserPassword(user *User, id int) (err error) {
	result, err := s.Repository.FindOneByEmail(user.Email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %s on database", user.Email)
		return errors.New(errMessage)
	}

	result.Password = user.Password

	err = s.Repository.Create(&result)
	return
}

func (s Service) DeleteUser(id int) error {
	err := s.Repository.Delete(id)

	if err != nil {
		return errors.New("user does not exist")
	}

	return nil
}

func (s Service) GetUserByEmail(email string) (user User, err error) {
	user, err = s.Repository.FindOneByEmail(email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %s on database", email)
		return User{}, errors.New(errMessage)
	}

	return
}

func (s Service) GetUserById(id int) (user User, err error) {
	user, err = s.Repository.FindById(id)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %d on database", id)
		return User{}, errors.New(errMessage)
	}

	return
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
