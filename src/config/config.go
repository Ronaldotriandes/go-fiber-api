package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		AppPort: getterEnv("APP_PORT", "3000"),
	}
}

func getterEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
