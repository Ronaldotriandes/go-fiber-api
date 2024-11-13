package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		AppPort: getterEnv("APP_PORT", "3000"),
		DBHost:  getterEnv("DB_HOST", "3000"),
		DBPort:  getterEnv("DB_PORT", "3000"),
		DBUser:  getterEnv("DB_USER", "3000"),
		DBPass:  getterEnv("DB_PASS", "3000"),
		DBName:  getterEnv("DB_NAME", "3000"),
	}
}

func getterEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
