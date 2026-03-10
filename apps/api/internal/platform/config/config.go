package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort 	 string
	DBHost 		 string
	DBPort 		 string
	DBUser 		 string
	DBPassword string
	DBName 		 string
	JWTSecret  string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	return Config{
		AppPort: 		getEnv("APP_PORT", "8080"),
		DBHost: 		getEnv("DB_HOST", "localhost"),
		DBPort: 		getEnv("DB_PORT", "5432"),
		DBUser: 		getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "Password123!"),
		DBName: 		getEnv("DB_NAME", "eventix"),
		JWTSecret: 	getEnv("JWT_SECRET", "my_secret_key"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
