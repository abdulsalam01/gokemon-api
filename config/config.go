package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file")
	}
}

func GetConfig() Configuration {

	InitEnv()

	configuration := Configuration{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	return configuration
}
