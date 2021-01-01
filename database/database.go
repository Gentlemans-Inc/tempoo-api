package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfig struct {
	host     string
	user     string
	password string
	name     string
	port     string
}

var Instance *gorm.DB

// ConnectDatabase creates the connection with postgres
func ConnectDatabase() {
	dbConfig := setupDatabase()
	p := dbConfig.port
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.host, port, dbConfig.user, dbConfig.password, dbConfig.name)
	Instance, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection opened to database")
	Instance.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")

	return
}

func setupDatabase() *databaseConfig {
	if os.Getenv("ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file!")
		}
	}

	return &databaseConfig{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name:     os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
	}
}
