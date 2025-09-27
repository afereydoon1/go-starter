package userservice

import "example.com/go-api/internal/domain/userentity"

type UserService interface {
	Register(user *userentity.User) (*userentity.User, error)
	Login(email, password string) (string, error)
}
