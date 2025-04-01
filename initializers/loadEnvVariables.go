package initializers

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}