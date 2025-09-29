// @title           Go Starter Kit API
// @version         1.0
// @description     Clean Architecture API in Golang with JWT Authentication and Category/User Management.

// @contact.name    Fereydoon Salemi
// @contact.email   afereydoon.s@gmail.com

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

// @host            localhost:3000
// @BasePath        /
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
	
	_ "example.com/go-api/docs" // Swagger generated docs
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	cfg := config.LoadConfig()
	database := db.Connect(cfg)
	utils.InitValidator()

	jwtService := utils.NewJWTService(cfg.JWTSecret)
	controllers := di.InitControllers(database, jwtService)

	r := gin.Default()
	router.RegisterRoutes(r, controllers)

	// serve Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	address := fmt.Sprintf("0.0.0.0:%s", cfg.AppPort)
	fmt.Printf("🚀 Server running at %s\n", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

