package main

import (
	"github.com/robbeheirman/todo/auth"
	"github.com/robbeheirman/todo/shared/app"
)

func GetRegisteredApps() []app.App {
	return []app.App{
		auth.NewApp(),
	}
}
