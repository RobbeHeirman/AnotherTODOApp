package api

import (
	"github.com/robbeheirman/todo/auth/persistence"
	"github.com/robbeheirman/todo/shared/routing"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserSchema struct {
	Name     string `json:"user"`
	Password string `json:"password"`
}

type Api struct {
	repository persistence.Repository
}

func NewApi(repository persistence.Repository) *Api {
	return &Api{repository}
}

func (api *Api) Register(user *UserSchema) (any, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password", err)
		return nil, &routing.RestError{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	userModel := persistence.User{
		user.Name,
		string(hashedPassword),
	}
	// TODO: Make specific saving errors
	err = api.repository.CreateUser(&userModel)
	if err != nil {
		return nil, &routing.RestError{
			Code:    http.StatusConflict,
			Message: "User already exists",
		}
	}
	return struct{}{}, nil
}
