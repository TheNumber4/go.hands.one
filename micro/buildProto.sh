#!/bin/sh
# prereq: brew install protobuf
go get -v github.com/golang/protobuf/protoc-gen-go && go get -v github.com/micro/protoc-gen-micro 

protoc -I$GOPATH/src --go_out=plugins=micro:${GOPATH}/src $(pwd)/proto/holiday/holiday.proto
