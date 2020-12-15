package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/api/config"
	"github.com/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB

// ConnectDatabase creates the connection with postgres
func ConnectDatabase() {
	var err error
	p := config.Env("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection opened to database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")
}
