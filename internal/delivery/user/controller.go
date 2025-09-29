package user

import (
	"net/http"

	"example.com/go-api/internal/usecase/userservice"
	"example.com/go-api/internal/delivery/user/requests"
	"example.com/go-api/internal/domain/userentity"
	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related endpoints
type UserHandler struct {
	userService userservice.UserService
}

func NewUserController(userService userservice.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Users
// @Accept json
// @Produce json
// @Param registerRequest body requests.RegisterRequest true "Register Request"
// @Success 201 {object} userentity.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	user := &userentity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := h.userService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// LoginUser godoc
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param loginRequest body requests.LoginRequest true "Login Request"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{Token: token})
}
