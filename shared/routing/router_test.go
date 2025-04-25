package routing

import (
	"net/http"
	"net/http/httptest"
	testing2 "shared/testing"
	"testing"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test"))
	if err != nil {
		panic(err)
	}
}

func runRequest(router *Router, request *http.Request) *http.Response {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w.Result()
}

func TestServeRoute(t *testing.T) {
	router := NewRouter()
	router.UseMiddleware(RedirectSlashes)
	router.HandleFunc("/test", getHandler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	result := runRequest(router, req)
	testing2.AssertEqual(t, http.StatusOK, result.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/test/", nil)
	result = runRequest(router, req)
	testing2.AssertEqual(t, http.StatusOK, result.StatusCode)
}

func TestSubRouter(t *testing.T) {
	router := NewRouter()
	router.UseMiddleware(RedirectSlashes)

	subRouter := NewRouter()
	subRouter.HandleFunc("/suffix", getHandler)

	router.Handle("/prefix", subRouter)
	req := httptest.NewRequest(http.MethodGet, "/prefix/suffix", nil)
	result := runRequest(router, req)
	testing2.AssertEqual(t, http.StatusOK, result.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/prefix/suffix/", nil)
	result = runRequest(router, req)
	testing2.AssertEqual(t, http.StatusOK, result.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/prefix", nil)
	result = runRequest(router, req)
	testing2.AssertEqual(t, http.StatusNotFound, result.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/prefix/", nil)
	result = runRequest(router, req)
	testing2.AssertEqual(t, http.StatusNotFound, result.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/suffix", nil)
	result = runRequest(router, req)
	testing2.AssertEqual(t, http.StatusNotFound, result.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/suffix/", nil)
	result = runRequest(router, req)
	testing2.AssertEqual(t, http.StatusNotFound, result.StatusCode)

}
