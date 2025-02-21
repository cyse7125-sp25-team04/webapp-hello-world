package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

var envConfig = fetchEnv()

func fetchEnv() Config {
	_ = godotenv.Load()

	return Config{
		DB_USERNAME: GetEnv("DB_USERNAME"),
		DB_PASSWORD: GetEnv("DB_PASSWORD"),
		DB_HOST:     GetEnv("DB_HOST"),
		DB_PORT:     GetEnv("DB_PORT"),
		DB_NAME:     GetEnv("DB_NAME"),
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func GetEnvConfig() Config {
	return envConfig
}
