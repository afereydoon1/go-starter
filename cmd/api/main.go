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
	// 1️⃣ Connect to database
	// ------------------------
	config.ConnectDB()
	fmt.Println("✅ Database connected!")

	// ------------------------
	// 2️⃣ Initialize Validator
	// ------------------------
	utils.InitValidator()
	fmt.Println("✅ Validator initialized!")

	// ------------------------
	// 3️⃣ Setup Gin router
	// ------------------------
	r := gin.Default()
	routes.RegisterRoutes(r)

	// ------------------------
	// 4️⃣ Start server
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
