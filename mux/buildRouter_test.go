package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type testRequest struct {
	route  string
	method string
	match  bool
}

func TestBuildRouter(t *testing.T) {
	router := buildRouter()

	expectedRoutes := []testRequest{
		{
			route:  "/",
			method: "GET",
			match:  true,
		},
	}

	for _, r := range expectedRoutes {
		match := &mux.RouteMatch{}

		if router.Match(httptest.NewRequest(r.method, r.route, nil), match) != r.match {
			t.Fatalf("Route %s/%s didn't return expected results !", r.method, r.route)
		}
	}
}
