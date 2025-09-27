package router

import (
	"github.com/gin-gonic/gin"
	"example.com/go-api/internal/di"
)

func RegisterRoutes(r *gin.Engine, controllers *di.AppControllers) {
	
	// -------------------
	// Users (Auth)
	// -------------------
	user := r.Group("/users")
	{
		user.POST("/register", controllers.UserController.Register)
		user.POST("/login", controllers.UserController.Login)
	}

	// -------------------------------------------------
	// Categories this routes can protect by middelware
	// -------------------------------------------------
	cat := r.Group("/categories")
	{
		cat.POST("", controllers.CategoryController.CreateCategory)
		cat.GET("/", controllers.CategoryController.ListCategories)
		cat.GET("/:id", controllers.CategoryController.GetCategory)
		cat.PUT("/:id", controllers.CategoryController.UpdateCategory)
		cat.DELETE("/:id", controllers.CategoryController.DeleteCategory)
	}

}
