package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnvVars() {
	if os.Getenv("ENV") != "PRODUCTION" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file!")
		}
	}
}
