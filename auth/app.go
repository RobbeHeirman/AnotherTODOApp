package auth

import (
	"github.com/robbeheirman/todo/auth/logic"
	"github.com/robbeheirman/todo/auth/persistence"
	"github.com/robbeheirman/todo/shared/routing"
	"net/http"
)

type App struct {
	repo        persistence.Repository
	signKey     string
	validateKey string
}

func NewApp(repo persistence.Repository, signKey string, validationKey string) *App {
	return &App{
		repo:        repo,
		signKey:     signKey,
		validateKey: validationKey,
	}
}

func (app *App) GetRouter() http.Handler {
	router := routing.NewRouter()
	newApi := logic.NewApi(app.repo, app.signKey)
	router.HandleFunc("/register", routing.RestPostHandleFunc(newApi.Register))
	return router
}

func (app *App) GetName() string { return "auth" }

func (app *App) Install() error {
	// TODO: Catch error
	return app.repo.Install()
}
