package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Test case 1: With environment variables
	os.Setenv("PORT", "8080")
	os.Setenv("APP_ENV", "production")
	os.Setenv("DB_CONNECTION", "mysql")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PASSWORD", "test_pass")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("JWT_SECRET", "test_secret")

	cfg := LoadConfig()

	assert.Equal(t, "8080", cfg.AppPort, "AppPort should match env variable")
	assert.Equal(t, "production", cfg.AppEnv, "AppEnv should match env variable")
	assert.Equal(t, "mysql", cfg.DBConnection, "DBConnection should match env variable")
	assert.Equal(t, "test_user", cfg.DBUser, "DBUser should match env variable")
	assert.Equal(t, "test_pass", cfg.DBPassword, "DBPassword should match env variable")
	assert.Equal(t, "localhost", cfg.DBHost, "DBHost should match env variable")
	assert.Equal(t, "3306", cfg.DBPort, "DBPort should match env variable")
	assert.Equal(t, "test_db", cfg.DBName, "DBName should match env variable")
	assert.Equal(t, "test_secret", cfg.JWTSecret, "JWTSecret should match env variable")

	// Test case 2: Without environment variables (fallback values)
	os.Clearenv()

	cfg = LoadConfig()

	assert.Equal(t, "3000", cfg.AppPort, "AppPort should use fallback")
	assert.Equal(t, "development", cfg.AppEnv, "AppEnv should use fallback")
	assert.Equal(t, "postgres", cfg.DBConnection, "DBConnection should use fallback")
	assert.Equal(t, "starter_user", cfg.DBUser, "DBUser should use fallback")
	assert.Equal(t, "starter_pass", cfg.DBPassword, "DBPassword should use fallback")
	assert.Equal(t, "db", cfg.DBHost, "DBHost should use fallback")
	assert.Equal(t, "5432", cfg.DBPort, "DBPort should use fallback")
	assert.Equal(t, "starter_db", cfg.DBName, "DBName should use fallback")
	assert.Equal(t, "your_default_secret", cfg.JWTSecret, "JWTSecret should use fallback")
}