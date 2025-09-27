package categoryservice

import (
    "example.com/go-api/internal/domain/categoryentity"
)

type categoryService struct {
    repo CategoryRepository
}

type CategoryRepository interface {
    Save(category *categoryentity.Category) error
    FindByID(id uint) (*categoryentity.Category, error)
    Update(category *categoryentity.Category) error
    Delete(id uint) error
    FindAll() ([]*categoryentity.Category, error)
}

func NewCategoryService(repo CategoryRepository) CategoryService {
    return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(c *categoryentity.Category) error {
    return s.repo.Save(c)
}

func (s *categoryService) GetCategoryByID(id uint) (*categoryentity.Category, error) {
    return s.repo.FindByID(id)
}

func (s *categoryService) UpdateCategory(c *categoryentity.Category) error {
    return s.repo.Update(c)
}

func (s *categoryService) DeleteCategory(id uint) error {
    return s.repo.Delete(id)
}

func (s *categoryService) ListCategories() ([]*categoryentity.Category, error) {
    return s.repo.FindAll()
}
