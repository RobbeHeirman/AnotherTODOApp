package routing

import (
	testing2 "github.com/robbeheirman/todo/shared/testing"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectSlashes(t *testing.T) {
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	redirectSlashes := RedirectSlashes(finalHandler)
	tests := []struct {
		input    string
		expected string
	}{
		{"/", "/"},
		{"/test", "/test/"},
		{"/test", "/test/"},
		{"/prefix/suffix", "/prefix/suffix/"},
		{"/prefix/suffix/", "/prefix/suffix/"},
	}
	for _, test := range tests {
		req := httptest.NewRequest("GET", test.input, nil)
		rr := httptest.NewRecorder()
		redirectSlashes.ServeHTTP(rr, req)
		testing2.AssertEqual(t, test.expected, req.URL.Path)
	}
}
