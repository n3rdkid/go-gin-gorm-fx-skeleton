package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//LoadEnv -> Load values from .env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load :: .env")
	}
}

// GetEnvByKey ->
func GetEnvByKey(key string) string {
	return os.Getenv(key)
}
