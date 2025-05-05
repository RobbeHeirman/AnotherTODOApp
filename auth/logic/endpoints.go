package logic

import (
	"crypto/rsa"
	"github.com/robbeheirman/todo/auth/models"
	"github.com/robbeheirman/todo/auth/persistence"
	"github.com/robbeheirman/todo/shared/routing"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"time"
)

type Api struct {
	repository persistence.Repository
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewApi(repository persistence.Repository, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *Api {
	return &Api{
		repository: repository,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (api *Api) Register(user *models.User) (any, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error hashing password", err)
		return nil, &routing.RestError{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	userModel := models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	// TODO: Make specific saving errors
	userId, err := api.repository.CreateUser(&userModel)
	if err != nil {
		slog.Error("Error creating user " + err.Error())
		return nil, &routing.RestError{
			Code:    http.StatusConflict,
			Message: "User already exists",
		}
	}
	jwt, err := CreateJwt(api.privateKey, userId.Id, time.Duration(72)*time.Hour)
	if err != nil {
		slog.Error("Error creating JWT", err)
		return nil, &routing.RestError{
			Code: http.StatusInternalServerError,
		}
	}
	return struct{ jwt string }{jwt}, nil
}

//func (api *Api) Login(user *models.User) (any, error) {
//	return nil, nil
//
//}
