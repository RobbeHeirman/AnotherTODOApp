package routing

import "net/http"

type Middleware func(http.Handler) http.Handler

type Router struct {
	mux        *http.ServeMux
	middleware []Middleware
}

func NewRouter() *Router {
	return &Router{
		mux:        http.NewServeMux(),
		middleware: make([]Middleware, 0),
	}
}

func (router *Router) UseMiddleware(middleware ...Middleware) *Router {
	router.middleware = append(router.middleware, middleware...)
	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := router.applyMiddleware(router.mux)
	handler.ServeHTTP(w, r)
}

func (router *Router) Handle(pattern string, handler http.Handler) *Router {
	router.mux.Handle(pattern+"/", http.StripPrefix(pattern, handler))
	return router
}

func (router *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) *Router {
	router.mux.Handle(pattern+"/", http.StripPrefix(pattern, handlerFunc))
	return router
}

func (router *Router) applyMiddleware(next http.Handler) http.Handler {
	for _, middleware := range router.middleware {
		next = middleware(next)
	}
	return next
}
