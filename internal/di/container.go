package di

import (
	"example.com/go-api/internal/delivery/user"
	"example.com/go-api/internal/delivery/category"
	"example.com/go-api/internal/infrastructure/db"
	"example.com/go-api/internal/usecase/categoryservice"
	"example.com/go-api/internal/usecase/userservice"
	"example.com/go-api/pkg/utils"
	"gorm.io/gorm"
)

type AppControllers struct {
	UserController     *user.UserHandler
	CategoryController *category.CategoryHandler
}

func InitControllers(database *gorm.DB, jwtService *utils.JWTService) *AppControllers {

	// -------------------
	// User (Auth)
	// -------------------
	userService := userservice.NewUserService(database, jwtService)
	userCtrl := user.NewUserController(userService)

	// -------------------
	// Category
	// -------------------
	categoryRepo := db.NewCategoryRepository(database)
	categoryService := categoryservice.NewCategoryService(categoryRepo)
	categoryCtrl := category.NewCategoryHandler(categoryService)

	return &AppControllers{
		CategoryController: categoryCtrl,
		UserController:     userCtrl,
	}
}
