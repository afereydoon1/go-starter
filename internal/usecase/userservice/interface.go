package userservice

import "example.com/go-api/internal/domain/userentity"

type UserRepository interface {
	Create(user *userentity.User) error
	FindByEmail(email string) (*userentity.User, error)
}
