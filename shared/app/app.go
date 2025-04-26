package app

import "shared/routing"

type App interface {
	GetRouter() *routing.Router
}
