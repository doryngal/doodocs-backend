package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	SMTPHost string
	SMTPPort string
	SMTPUser string
	SMTPPass string
}

func LoadConfig(filePath string) (*Config, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		Port:     getEnv("PORT", "8080"),
		SMTPHost: getEnv("SMTP_HOST", ""),
		SMTPPort: getEnv("SMTP_PORT", ""),
		SMTPUser: getEnv("SMTP_USER", ""),
		SMTPPass: getEnv("SMTP_PASS", ""),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
