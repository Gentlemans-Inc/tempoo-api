package user

import (
	"github.com/api/internal/database"
	"gorm.io/gorm"
)

// UserRepository Contract
type UserRepository interface {
	FindAll() (users *gorm.DB, err error)
	FindOneByEmail(email string) (user User, err error)
	FindById(id int) (user User, err error)
	Create(user *User) error
	Delete(id int) error
}

//NewUserRepository repository postgres implementation
func NewUserRepository() UserRepository {
	return &Repository{
		DB: database.Instance,
	}
}


