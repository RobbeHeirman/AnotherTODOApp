package main

import (
	"github.com/robbeheirman/todo/shared/routing"
	"log"
	"net/http"
)

func main() {
	mainRouter := routing.NewRouter()
	mainRouter.UseMiddleware(routing.RedirectSlashes)
	for _, app := range GetRegisteredApps() {
		mainRouter.Handle("/"+app.GetName(), app.GetRouter())
	}

	if err := http.ListenAndServe(":8080", mainRouter); err != nil {
		log.Fatal(err)
	}
}
