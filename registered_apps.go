package main

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/robbeheirman/todo/auth"
	"github.com/robbeheirman/todo/auth/persistence/postgres"
	"github.com/robbeheirman/todo/shared/app"
	"log"
	"log/slog"
	"os"
	"strconv"
)

func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		slog.Error("Error generating rsa key", err)
		panic(err)
	}
	publicKey := privateKey.PublicKey
	return privateKey, &publicKey
}

func GetRegisteredApps() ([]app.App, error) {
	DbUsername := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbDatabase := os.Getenv("DB_DATABASE")

	signKey, validationKey := generateKeys()
	DbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return []app.App{
		auth.NewApp(postgres.NewRepository(DbHost, DbPort, DbDatabase, DbUsername, DbPassword), signKey, validationKey),
	}, nil
}
