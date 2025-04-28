package app

import (
	"net/http"
)

type App interface {
	GetRouter() http.Handler
	GetName() string
	Install() error
}
