package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	buildRouter().ServeHTTP(w, r)

	if w.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Expected json but got %s", w.Header().Get("Content-Type"))
	}
	if w.Code != http.StatusOK {
		t.Fatalf("Expected code %d but got %d",
			http.StatusOK,
			w.Code)
	}

	s := string(w.Body.Bytes())
	if !strings.Contains(s, "root") {
		t.Fatalf("Body does not contains root => %s", s)
	}
}
