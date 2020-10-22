#!/bin/sh
go env -w CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" 
go env -w GOPROXY=https://goproxy.cn,direct

# pkill -9 daigouadmin

file="bin"
if [ -d $file ];then
  echo "删除$file"
  rm -rf $file
fi

cd code 
go mod tidy
# go get -v -u github.com/swaggo/swag/cmd/swag
# swag init
cp -r bin.template/ ../bin
go build -o ../bin/daigouadmin -tags=jsoniter main.go
cd ../bin
chmod +x ./daigouadmin
./daigouadmin

# nohup ./daigouadmin >> nohup.log 2>&1 &