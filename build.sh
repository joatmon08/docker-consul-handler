#!/bin/bash

GOOS=linux GOARCH=amd64 go build handler.go
docker build -t consul-handler:latest .