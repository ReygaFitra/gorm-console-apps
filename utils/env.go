package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConfigEnv(key string) string {

	err := godotenv.Load("config.env"); if err != nil {
		log.Fatalln("error load .env file")
	}

	return os.Getenv(key)
}