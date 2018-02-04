package main

import (
	"fmt"
	"net/http"
)

func main() {
	bind, port := parse()
	addr := fmt.Sprintf("%s:%d",
		*bind,
		*port)

	http.HandleFunc("/", handle)
	fmt.Printf("Server started on %s\n", addr)
	http.ListenAndServe(
		addr,
		nil)
}
