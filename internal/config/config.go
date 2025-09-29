package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	ApiKey     string
	ApiURL     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ApiKey:     getEnv("API_KEY", ""),
		ApiURL:     getEnv("API_URL", "https://api.openweathermap.org"),
	}, nil
}

func getEnv(key, defoultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defoultValue
}
