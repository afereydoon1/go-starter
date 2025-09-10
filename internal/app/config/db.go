package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/go-api/internal/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbType := strings.ToLower(os.Getenv("DB_CONNECTION"))
	var dsn string
	var dialector gorm.Dialector

	switch dbType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		dialector = mysql.Open(dsn)

	case "postgres", "pgsql":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
		dialector = postgres.Open(dsn)

	default:
		log.Fatalf("❌ Unsupported DB_CONNECTION: %s", dbType)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// ✅ Auto migrate all models here
	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	DB = db
	log.Println("✅ Connected and migrated DB with", dbType)
}
