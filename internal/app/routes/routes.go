package routes

import (
	"example.com/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.CreateUser)

	// Protected routes

	users := r.Group("/users")
	{
		users.POST("", controllers.CreateUser)
		users.GET("", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
