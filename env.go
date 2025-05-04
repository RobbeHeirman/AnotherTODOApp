//go:build dev

package main

import (
	_ "embed"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//go:embed private_key.pem
var privateKey string

//go:embed public_key.pem
var publicKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	} else {
		log.Println("Loading .env file")
	}

	err := os.Setenv("SIGNING_KEY", privateKey)
	if err != nil {
		panic(err)
	}

	err = os.Setenv("VALIDATION_KEY", publicKey)
	if err != nil {
		panic(err)
	}
}
