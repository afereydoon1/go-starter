package categoryservice

import "example.com/go-api/internal/domain/categoryentity"

type CategoryService interface {
	CreateCategory(category *categoryentity.Category) error
	GetCategoryByID(id uint) (*categoryentity.Category, error)
	UpdateCategory(category *categoryentity.Category) error
	DeleteCategory(id uint) error
	ListCategories() ([]*categoryentity.Category, error)
}