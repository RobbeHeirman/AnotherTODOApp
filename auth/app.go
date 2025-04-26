package auth

import (
	"github.com/robbeheirman/todo/auth/api"
	"github.com/robbeheirman/todo/auth/persistence"
	"github.com/robbeheirman/todo/shared/routing"
	"net/http"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) GetRouter() http.Handler {
	router := routing.NewRouter()
	repository := persistence.NewPostgresRepository()
	newApi := api.NewApi(repository)
	router.HandleFunc("/register", routing.RestPostHandleFunc(newApi.Register))
	return router
}

func (app *App) GetName() string { return "auth" }
