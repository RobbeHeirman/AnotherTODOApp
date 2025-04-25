package testing

import (
	"fmt"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, expected T, result T) {
	t.Helper()
	if expected != result {
		t.Errorf(fmt.Sprintf("Expected %v, got %v", expected, result))
	}
}
