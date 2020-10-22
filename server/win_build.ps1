$Env:GOOS="windows" 
$Env:GOARCH="amd64"
$Env:CGO_ENABLED=0

if (Test-Path bin/server.exe){(Remove-Item bin/server.exe)}

go build -o bin/server.exe -tags=jsoniter main.go