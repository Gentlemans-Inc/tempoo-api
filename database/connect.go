package database

import (
	"github.com/api/config"
	"github.com/api/models"	
	"fmt"
	"strconv"
	
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// ConnectDatabase creates the connection with postgres
func ConnectDatabase() {
	var err error
	p := config.Env("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Env("DB_HOST"), port, config.Env("DB_USER"), config.Env("DB_PASSWORD"), config.Env("DB_NAME"))
	
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("Connection opened to database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")
}