# simple.web
Simple Web demonstrate a very simple web service that returns a quote and author in JSON.

## params.go
This file contains the parse function. It is used to parse command line parameters for the bind (what address listen) and what port should be used.
It demonstrates 

* flag package: How to parse command line easely

## handler.go
This file contains the handling code. It is used to create a reply to the client.

It demonstrates
* structure in go
* annotations (JSON:)
* http package ( response and request)
* json encoder
* structure creation in go

## server.go
This file contains the actual server startup code.

It demonstrates
* http handleFunc/server

## Dockerfile
This file run a docker go machine to build project code and after build, create a run docker image (alpine)

It demonstrates
* multi stage docker image and go
* Golang image size (11mb)

