package main

import (
	"fmt"
	"log"

	"example.com/go-api/internal/config"
	"example.com/go-api/internal/di"
	"example.com/go-api/internal/infrastructure/db"
	"example.com/go-api/internal/router"
	"example.com/go-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database := db.Connect(cfg)
	utils.InitValidator()

	jwtService := utils.NewJWTService(cfg.JWTSecret)
	controllers := di.InitControllers(database, jwtService)

	r := gin.Default()
	router.RegisterRoutes(r, controllers)

	address := fmt.Sprintf("0.0.0.0:%s", cfg.AppPort)
	fmt.Printf("🚀 Server running at %s\n", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
