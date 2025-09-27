package config

import (
	"log"
	"os"
)

type Config struct {
	AppPort      string
	AppEnv       string
	DBConnection string
	DBUser       string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	JWTSecret    string
}

// LoadConfig returns a Config instance instead of assigning to a global variable
func LoadConfig() *Config {
	cfg := &Config{
		AppPort:      getEnv("PORT", "3000"),
		AppEnv:       getEnv("APP_ENV", "development"),
		DBConnection: getEnv("DB_CONNECTION", "postgres"),
		DBUser:       getEnv("DB_USER", "starter_user"),
		DBPassword:   getEnv("DB_PASSWORD", "starter_pass"),
		DBHost:       getEnv("DB_HOST", "db"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBName:       getEnv("DB_NAME", "starter_db"),
		JWTSecret:    getEnv("JWT_SECRET", "your_default_secret"),
	}

	log.Println("✅ Config loaded successfully")
	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
