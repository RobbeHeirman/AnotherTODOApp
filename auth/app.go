package auth

import (
	"github.com/robbeheirman/todo/auth/logic"
	"github.com/robbeheirman/todo/auth/persistence"
	"github.com/robbeheirman/todo/shared/routing"
	"net/http"
)

type App struct {
	repo persistence.Repository
}

func NewApp(repo persistence.Repository) *App {
	return &App{
		repo: repo,
	}
}

func (app *App) GetRouter() http.Handler {
	router := routing.NewRouter()
	newApi := logic.NewApi(app.repo)
	router.HandleFunc("/register", routing.RestPostHandleFunc(newApi.Register))
	return router
}

func (app *App) GetName() string { return "auth" }

func (app *App) Install() error {
	// TODO: Catch error
	return app.repo.Install()
}
