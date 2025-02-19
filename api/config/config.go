package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	CFG *Config

	// Global Vars
	ZINC_URL_API          string
	ZINC_GET_DOCUMENT_URL string
	ZINC_SEARCH_URL       string
	IS_DEVELOPMENT        bool = false
)

type Config struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	ServerPort       string
	AllowedOrigins   string
}

func LoadConfig() {
	loadEnv()
}

func loadEnv() {
	if os.Getenv("APP_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			panic(".env not found")
		}
	}

	CFG = &Config{
		DatabaseUser:     getEnv("DATABASE_USER"),
		DatabasePassword: getEnv("DATABASE_PASSWORD"),
		DatabaseHost:     getEnv("DATABASE_HOST"),
		DatabasePort:     getEnv("DATABASE_PORT"),
		DatabaseName:     getEnv("DATABASE_NAME"),
		ServerPort:       getEnv("SERVER_PORT"),
		AllowedOrigins:   getEnv("ALLOWED_ORIGINS"),
	}

	IS_DEVELOPMENT = os.Getenv("APP_MODE") == "development"

	ZINC_URL_API = CFG.DatabaseHost + ":" + CFG.DatabasePort
	ZINC_SEARCH_URL = fmt.Sprintf("%v/es/%v/_search", ZINC_URL_API, CFG.DatabaseName)
	ZINC_GET_DOCUMENT_URL = fmt.Sprintf("%v/api/%v/_doc", ZINC_URL_API, CFG.DatabaseName)
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		// fmt.Printf("Environment variable %v is set. %v\n", key, value)
		return value
	}

	fmt.Printf("Environment variable %v is not set.\n", key)
	panic("Environment variable " + key + " is not set.\n")
}
