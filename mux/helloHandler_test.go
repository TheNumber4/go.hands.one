package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/hello/Jim", nil)
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
	if !strings.Contains(s, "Jim") {
		t.Fatalf("Body does not contains Jim ! ==> %s", s)
	}
}
