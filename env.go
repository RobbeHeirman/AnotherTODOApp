//go:build dev

package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err)
	} else {
		log.Println("Loading .env file")
	}
}
