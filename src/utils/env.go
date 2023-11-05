package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load()
	for _, envVar := range EnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatal(envVar + " is not seted in .env file")
		}
	}
}
