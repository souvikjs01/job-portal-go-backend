package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("error loading env file")
	// }

	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found — falling back to system environment variables")
		} else {
			log.Println(".env file loaded successfully")
		}
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
