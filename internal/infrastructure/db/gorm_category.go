package db

import (
    "gorm.io/gorm"
    "example.com/go-api/internal/domain/categoryentity"
)

type CategoryGormRepo struct {
    DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryGormRepo {
    return &CategoryGormRepo{DB: db}
}

func (r *CategoryGormRepo) Save(c *categoryentity.Category) error {
    return r.DB.Create(c).Error
}

func (r *CategoryGormRepo) FindByID(id uint) (*categoryentity.Category, error) {
    var category categoryentity.Category
    err := r.DB.First(&category, id).Error
    return &category, err
}

func (r *CategoryGormRepo) Update(c *categoryentity.Category) error {
    return r.DB.Save(c).Error
}

func (r *CategoryGormRepo) Delete(id uint) error {
    return r.DB.Delete(&categoryentity.Category{}, id).Error
}

func (r *CategoryGormRepo) FindAll() ([]*categoryentity.Category, error) {
    var categories []*categoryentity.Category
    err := r.DB.Find(&categories).Error
    return categories, err
}
