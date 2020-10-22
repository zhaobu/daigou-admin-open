#!/bin/sh

go env -w CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" 
go env -w GOPROXY=https://goproxy.cn,direct

rm bin/daigouadmin
swag init
go build -o bin/daigouadmin main.go