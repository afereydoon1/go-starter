package category

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "example.com/go-api/internal/domain/categoryentity"
    "example.com/go-api/internal/delivery/category/requests"
    "example.com/go-api/internal/usecase/categoryservice"
)

type CategoryHandler struct {
    service categoryservice.CategoryService
}

func NewCategoryHandler(service categoryservice.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    var req requests.CreateCategoryRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category := &categoryentity.Category{
        Name: req.Name,
        Slug: req.Slug,
    }

    if err := h.service.CreateCategory(category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    category, err := h.service.GetCategoryByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
    var req requests.UpdateCategoryRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, _ := strconv.Atoi(c.Param("id"))
    category := &categoryentity.Category{
        ID:   uint(id),
        Name: req.Name,
        Slug: req.Slug,
    }

    if err := h.service.UpdateCategory(category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.service.DeleteCategory(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
    categories, err := h.service.ListCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}
