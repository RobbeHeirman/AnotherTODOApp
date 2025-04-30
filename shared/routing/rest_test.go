package routing

import (
	"encoding/json"
	testing2 "github.com/robbeheirman/todo/shared/testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestData struct {
	Dummy string `json:"dummy"`
}

type TestResult struct {
	Result string `json:"result"`
}

func postFunc(data *TestData, w http.ResponseWriter, r *http.Request) *TestResult {
	return &TestResult{
		Result: data.Dummy + "Result",
	}
}

func TestPostWrapper(t *testing.T) {
	handler := RestPostHandleFunc(postFunc)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"dummy\": \"test\"}"))
	handler.ServeHTTP(w, r)
	expected := TestResult{
		Result: "testResult",
	}

	var result TestResult
	err := json.NewDecoder(w.Body).Decode(&result)
	if err != nil {
		t.Error("Other error:", err)
	}
	testing2.AssertEqual(t, expected, result)
}
