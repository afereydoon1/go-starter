package db

import (
	"fmt"
	"log"
	"strings"

	"example.com/go-api/internal/config"
	"example.com/go-api/internal/domain/categoryentity"
	"example.com/go-api/internal/domain/userentity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dbType := strings.ToLower(cfg.DBConnection)
	var dsn string
	var dialector gorm.Dialector

	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dsn)
	case "postgres", "pgsql":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		dialector = postgres.Open(dsn)
	default:
		log.Printf("❌ Unsupported DB_CONNECTION: %s", dbType)
	    panic("Unsupported DB_CONNECTION")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto migrate domain models
	if err := db.AutoMigrate(&userentity.User{},&categoryentity.Category{}); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	log.Println("✅ Connected and migrated DB with", dbType)
	return db
}
