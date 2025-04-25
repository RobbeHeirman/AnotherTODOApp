package routing

import "net/http"

func RedirectSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path != "/" && !(path[len(path)-1] == '/') {
			r.URL.Path += "/"
		}
		next.ServeHTTP(w, r)
	})
}
