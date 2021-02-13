package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config type represents the db connection string
type Config struct {
	host     string
	user     string
	password string
	name     string
	port     string
	sslMode  string
}

// Instance of the gorm connection
var Instance *gorm.DB

// ConnectDatabase creates the connection with postgres
func ConnectDatabase() {
	dbConfig := SetupDatabase()
	p := dbConfig.port
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbConfig.host, port, dbConfig.user, dbConfig.password, dbConfig.name, dbConfig.sslMode)
	Instance, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return
}

// SetupDatabase returns an databaseConfig pointer
func SetupDatabase() *Config {
	return &Config{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name:     os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		sslMode:  os.Getenv("DB_SSLMODE"),
	}
}
