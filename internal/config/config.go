package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, reading from system env")
	}
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}