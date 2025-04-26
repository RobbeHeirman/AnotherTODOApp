package auth

import "github.com/robbeheirman/todo/shared/routing"

type App struct{}

func (app *App) GetRouter() *routing.Router {
	router := routing.NewRouter()

	return router
}

func NewApp() *App {
	return &App{}
}
