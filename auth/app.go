package auth

import (
	"crypto/rsa"
	"github.com/robbeheirman/todo/auth/logic"
	"github.com/robbeheirman/todo/auth/persistence"
	"github.com/robbeheirman/todo/shared/routing"
	"net/http"
)

type App struct {
	repo        persistence.Repository
	signKey     *rsa.PrivateKey
	validateKey *rsa.PublicKey
}

func NewApp(repo persistence.Repository, signKey *rsa.PrivateKey, validationKey *rsa.PublicKey) *App {
	return &App{
		repo:        repo,
		signKey:     signKey,
		validateKey: validationKey,
	}
}

func (app *App) GetRouter() http.Handler {
	router := routing.NewRouter()
	newApi := logic.NewApi(app.repo, app.signKey, app.validateKey)
	router.HandleFunc("/register", routing.RestPostHandleFunc(newApi.Register))
	return router
}

func (app *App) GetName() string { return "auth" }

func (app *App) Install() error {
	// TODO: Catch error
	return app.repo.Install()
}
