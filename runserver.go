package main

import (
	"github.com/robbeheirman/todo/shared/routing"
	"log"
	"net/http"
)

func main() {
	mainRouter := routing.NewRouter()
	mainRouter.UseMiddleware(routing.RedirectSlashes)
	apps, err := GetRegisteredApps()
	if err != nil {
		log.Fatal("Fatal Error: ", err)
	}

	for _, app := range apps {
		err := app.Install()
		if err != nil {
			log.Println("App Did not install:", err)
		}
	}

	for _, app := range apps {
		mainRouter.Handle("/"+app.GetName(), app.GetRouter())
	}

	if err := http.ListenAndServe(":8081", mainRouter); err != nil {
		log.Fatal(err)
	}
}
