package repository

import (
	"fmt"
	"github.com/api/database"
	"github.com/api/models"
	"github.com/api/repository/user"
	"gorm.io/gorm"
)

// UserRepository Contract
type UserRepository interface {
	FindAll() (users *gorm.DB, err error)
	FindOneByEmail(email string) (user *gorm.DB, err error)
	Create(user *models.User) error
}

//NewUserRepository repository postgres implementation
func NewUserRepository() UserRepository {
	fmt.Println(database.Instance)
	return &user.Repository{
		DB: database.Instance,
	}
}
