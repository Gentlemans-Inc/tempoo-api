package models

import (
	"gorm.io/gorm"
)

// User is the model for the DB
type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Name     string `json:"name"`
	Password string `gorm:"not null" json:"password"`
}

// UserLogin handles data for /auth/login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser handles data for POST in /user
type NewUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UserDetails handles data for PUT/POST in /users
type UserDetails struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
