package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

var Appcfg *Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using system env")
	}

	Appcfg = &Config{
		Port: getEnv("PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "upi_app"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}


func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)

	if !exists {
		return fallback;
	}

	return val
}