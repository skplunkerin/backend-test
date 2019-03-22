#!/bin/sh

GOOS=darwin GOARCH=386 go build -o build/imgsrv-darwin-386 main.go
GOOS=darwin GOARCH=amd64 go build -o build/imgsrv-darwin-amd64 main.go

GOOS=linux GOARCH=386 go build -o build/imgsrv-linux-386 main.go
GOOS=linux GOARCH=amd64 go build -o build/imgsrv-linux-amd64 main.go

GOOS=windows GOARCH=386 go build -o build/imgsrv-windows-386 main.go
GOOS=windows GOARCH=amd64 go build -o build/imgsrv-windows-amd64 main.go
