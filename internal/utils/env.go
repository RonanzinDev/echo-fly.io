package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ronanzindev/echo-fly.io/internal/config"
)

func LoadEnv() {
	godotenv.Load()
	for _, envVar := range config.EnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatal(envVar + " is not seted in .env file")
		}
	}
}
