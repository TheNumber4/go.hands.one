#!/bin/sh
protoc -I$GOPATH/src --go_out=plugins=micro:${GOPATH}/src $(pwd)/proto/holiday/holiday.proto
