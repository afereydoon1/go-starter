package db

import (
    "os"
    "testing"

    "example.com/go-api/internal/config"
    "example.com/go-api/internal/domain/userentity"
    "github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
    cfg := &config.Config{
        DBConnection: "postgres",
        DBHost:       os.Getenv("DB_HOST"),      
        DBUser:       os.Getenv("DB_USER"),
        DBPassword:   os.Getenv("DB_PASSWORD"),
        DBName:       os.Getenv("DB_NAME"),
        DBPort:       os.Getenv("DB_PORT"),
        JWTSecret:    os.Getenv("JWT_SECRET"),   
    }

    if cfg.DBHost == "" {
        cfg.DBHost = "localhost"
    }
    if cfg.DBUser == "" {
        cfg.DBUser = "test_user"
    }
    if cfg.DBPassword == "" {
        cfg.DBPassword = "test_pass"
    }
    if cfg.DBName == "" {
        cfg.DBName = "test_db"
    }
    if cfg.DBPort == "" {
        cfg.DBPort = "5432"
    }
    if cfg.JWTSecret == "" {
        cfg.JWTSecret = "test_secret"
    }

    t.Run("Successful connection to PostgreSQL", func(t *testing.T) {
        db := Connect(cfg)
        assert.NotNil(t, db, "Database connection should not be nil")

        var result int
        db.Raw("SELECT 1").Scan(&result)
        assert.Equal(t, 1, result, "Database should respond to query")

        err := db.AutoMigrate(&userentity.User{})
        assert.NoError(t, err, "AutoMigrate should not return error")
    })

    t.Run("Invalid DB type", func(t *testing.T) {
        cfgInvalid := &config.Config{
            DBConnection: "invalid",
            DBHost:       "localhost",
            DBUser:       "test_user",
            DBPassword:   "test_pass",
            DBName:       "test_db",
            DBPort:       "5432",
        }

        defer func() {
            if r := recover(); r == nil {
                t.Errorf("Expected panic for unsupported DB type")
            }
        }()
        Connect(cfgInvalid)
    })
}