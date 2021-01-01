package user

import (
	userError "github.com/api/helpers/errors/user"
	"github.com/api/models"
	"gorm.io/gorm"
	"log"
)

// Repository concrete type
type 	Repository struct {
	DB *gorm.DB // this can be any gorm instance
}

func (r Repository) FindAll() (users *gorm.DB, err error) {
	result := r.DB.Find(&users)
	err = result.Error
	if err != nil {
		log.Fatal("An error occurred in the query: ", err)
	}
	return
}

func (r Repository) FindOneByEmail(email string) (result *gorm.DB, err error) {
	aux := models.User{}
	result = r.DB.First(&aux, "email = ?", email)
	err = result.Error
	if result.RowsAffected > 0 {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return
}

func (r Repository) Create(user *models.User) error {
	result := r.DB.Create(user)
	err := result.Error
	rowsCount := result.RowsAffected
	if err != nil || rowsCount <= 0 {
		return userError.CannotCreateError{}
	}
	return nil
}
