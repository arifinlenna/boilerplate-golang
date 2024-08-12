package appconfig

import (
	"log"

	"github.com/joho/godotenv"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// config.DB = config.Oracle()
}
