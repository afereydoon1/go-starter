package main

import (
	"fmt"
	"log"
	"os"

	"example.com/go-api/internal/app/config"
	"example.com/go-api/internal/app/routes"
	"example.com/go-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// ------------------------
	// 1️⃣ Load environment variables
	// ------------------------
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		dbHost = "db" // default service name in docker-compose
	}

	// ------------------------
	// 2️⃣ Connect to database
	// ------------------------
	err := config.ConnectDB(dbUser, dbPass, dbName, dbHost)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	fmt.Println("✅ Database connected!")

	// ------------------------
	// 3️⃣ Initialize Validator
	// ------------------------
	utils.InitValidator()
	fmt.Println("✅ Validator initialized!")

	// ------------------------
	// 4️⃣ Setup Gin router
	// ------------------------
	r := gin.Default()
	routes.RegisterRoutes(r)

	// ------------------------
	// 5️⃣ Start server
	// ------------------------
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("🚀 Server running at %s\n", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
