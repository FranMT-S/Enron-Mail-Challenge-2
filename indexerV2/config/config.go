package config

import (
	"fmt"
	"log"
	"os"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/shared"
	"github.com/joho/godotenv"
)

var (
	CFG                   *Config
	IS_PROFILER           bool
	IS_PROFILER_TRACER    bool
	ZINC_URL_API          string
	ZINC_BULK_API_URL     string
	ZINC_CREATE_INDEX_URL string
	IS_DEVELOPMENT        bool = false
)

type Config struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
}

func LoadConfig() {
	shared.CreateFolderIfNotExist("./profiling")
	shared.CreateFolderIfNotExist("./logs")
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found\n")
	}

	CFG = &Config{
		DatabaseUser:     getEnv("DATABASE_USER"),
		DatabasePassword: getEnv("DATABASE_PASSWORD"),
		DatabaseHost:     getEnv("DATABASE_HOST"),
		DatabasePort:     getEnv("DATABASE_PORT"),
		DatabaseName:     getEnv("DATABASE_NAME"),
	}

	IS_DEVELOPMENT = os.Getenv("APP_ENV") == "development"
	IS_PROFILER = os.Getenv("PROFILER") == "true"
	IS_PROFILER_TRACER = os.Getenv("PROFILER_TRACE") == "true"

	ZINC_URL_API = CFG.DatabaseHost + ":" + CFG.DatabasePort + "/api"
	ZINC_BULK_API_URL = ZINC_URL_API + "/_bulkv2"
	ZINC_CREATE_INDEX_URL = ZINC_URL_API + "/index"
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		fmt.Printf("Environment variable %v is set. %v\n", key, value)
		return value
	}

	fmt.Printf("Environment variable %v is not set.\n", key)
	panic("Environment variable " + key + " is not set.\n")
}
