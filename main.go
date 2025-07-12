package main

import (
	"example.com/go-api/config"
	"example.com/go-api/routes"
	"example.com/go-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	utils.InitValidator()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
