package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type someReply struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(&someReply{
		Quote:  "FAKE NEWS!",
		Author: "Donald Trump",
	})

	if err != nil {
		panic(fmt.Errorf("Failed to convert reply: %e", err))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)

	fmt.Printf("Hanlded request from %s\n", r.RemoteAddr)
}
