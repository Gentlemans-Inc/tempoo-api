package repository

import (
	"errors"

	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user"
	"gorm.io/gorm"
)

// Repository concrete type
type Repository struct {
	DB *gorm.DB // this can be any gorm instance
}

// FindAll find the users in DB
func (r Repository) FindAll() (users []user.User, err error) {
	result := r.DB.Find(&users)
	err = result.Error
	return
}

func (r Repository) FindOneByEmail(email string) (user user.User, err error) {
	result := r.DB.First(&user, "email = ?", email)
	err = result.Error
	return
}

func (r Repository) FindById(id int) (user user.User, err error) {
	result := r.DB.First(&user, "id = ?", id)
	err = result.Error
	return
}

func (r Repository) Delete(id int) (err error) {
	result := r.DB.Delete(&user.User{}, "id = ?", id)
	err = result.Error
	return
}

func (r Repository) Create(user *user.User) error {
	result := r.DB.Create(user)
	err := result.Error
	rowsCount := result.RowsAffected
	if err != nil || rowsCount <= 0 {
		return errors.New("cannot create user")
	}
	return nil
}
