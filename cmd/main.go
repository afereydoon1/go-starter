package main

import (
	"example.com/go-api/internal/app/config"
	"example.com/go-api/internal/app/routes"
	"example.com/go-api/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	utils.InitValidator()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
