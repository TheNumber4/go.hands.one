#!/bin/sh

docker build -t web . && \
		docker run --rm -p 8080:8080 --name web web 
