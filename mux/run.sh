#!/bin/sh

docker build -t webmux . && \
		docker run --rm -p 8080:8080 --name webmux webmux
