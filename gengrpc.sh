#!/bin/sh

protoc --go_out=$GOPATH/src -I=./gcommand ./gcommand/gcommand.proto
protoc --go-grpc_out=$GOPATH/src -I=./gcommand ./gcommand/gcommand.proto
