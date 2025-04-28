package main

import (
	"github.com/robbeheirman/todo/auth"
	"github.com/robbeheirman/todo/auth/persistence/postgres"
	"github.com/robbeheirman/todo/shared/app"
	"log"
	"os"
	"strconv"
)

func GetRegisteredApps() ([]app.App, error) {
	DbUsername := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbDatabase := os.Getenv("DB_DATABASE")

	DbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return []app.App{
		auth.NewApp(postgres.NewRepository(DbHost, DbPort, DbDatabase, DbUsername, DbPassword)),
	}, nil
}
