package models

// User is the model for the DB
type User struct {
	Name     string
	Password string
	Email    string
}

// UserLogin handles data for /auth/login
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
