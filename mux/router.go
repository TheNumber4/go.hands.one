package main

import "github.com/gorilla/mux"

func buildRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/hello/{name}", helloHandler)

	return r
}
