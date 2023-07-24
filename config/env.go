package config

import (
	"log"

	"api.virak.me/utils"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if utils.IsNotNil(err) {
		log.Fatal("Error loading .env file")
	}
}
