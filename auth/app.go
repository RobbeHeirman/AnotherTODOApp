package auth

import "github.com/robbeheirman/todo/shared/routing"

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) GetRouter() *routing.Router {
	router := routing.NewRouter()
	router.HandleFunc("/register", routing.RestPostHandleFunc(Register))

	return router
}

func (app *App) GetName() string { return "auth" }
