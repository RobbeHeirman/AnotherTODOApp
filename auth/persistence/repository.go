package persistence

import "github.com/robbeheirman/todo/auth/models"

type Repository interface {
	Install() error
	CreateUser(user *models.User) error
}
