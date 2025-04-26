package app

import "github.com/robbeheirman/todo/shared/routing"

type App interface {
	GetRouter() *routing.Router
	GetName() string
}
