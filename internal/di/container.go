package di

import (
	"example.com/go-api/internal/delivery/category"
	"example.com/go-api/internal/infrastructure/db"
	"example.com/go-api/internal/usecase/categoryservice"
	"example.com/go-api/pkg/utils"
	"gorm.io/gorm"
)

type AppControllers struct {
	CategoryController *category.CategoryHandler
	
}

func InitControllers(database *gorm.DB, jwtService *utils.JWTService) *AppControllers {
	// Repository
	categoryRepo := db.NewCategoryRepository(database)

	// Service
	categoryService := categoryservice.NewCategoryService(categoryRepo)

	// Controller
	categoryCtrl := category.NewCategoryHandler(categoryService)

	return &AppControllers{
		CategoryController: categoryCtrl,
	}
}
