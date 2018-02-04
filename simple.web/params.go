package main

import "flag"

func parse() (*string, *int) {
	portPtr := flag.Int("port", 8080, "Port to use. Default to 8080")
	bindPtr := flag.String("host", "", "Bind address")

	flag.Parse()

	return bindPtr, portPtr
}
