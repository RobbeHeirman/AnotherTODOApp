package persistence

import (
	"github.com/robbeheirman/todo/auth/models"
)

type Repository interface {
	Install() error
	CreateUser(user *models.User) (*models.UserId, error)
	GetUserByEmail(User *models.User) (models.UserLogsInDb, error)
}
