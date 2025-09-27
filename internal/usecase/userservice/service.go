package userservice

import (
	"errors"
	"example.com/go-api/internal/domain/userentity"
	"example.com/go-api/pkg/utils"
	"gorm.io/gorm"
)

type userService struct {
	db         *gorm.DB
	jwtService *utils.JWTService
}

func NewUserService(db *gorm.DB, jwtService *utils.JWTService) UserService {
	return &userService{
		db:         db,
		jwtService: jwtService,
	}
}

func (s *userService) Register(user *userentity.User) (*userentity.User, error) {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	var user userentity.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := s.jwtService.GenerateJWT(user.ID) 
	if err != nil {
		return "", err
	}
	return token, nil
}




