package appconfig

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/lenna-ai/azureOneSmile.git/config"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file")
	}
	config.DB = config.Database()
}
