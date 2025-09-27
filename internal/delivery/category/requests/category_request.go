package requests

type CreateCategoryRequest struct {
    Name string `json:"name" binding:"required"`
    Slug string `json:"slug" binding:"required"`
}

type UpdateCategoryRequest struct {
    Name string `json:"name" binding:"required"`
    Slug string `json:"slug" binding:"required"`
}
