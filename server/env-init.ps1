# 复制模版文件
if (Test-Path .vscode) { (Remove-Item -r .vscode) }
if (Test-Path bin) { (Remove-Item -r bin) }

Copy-Item -r .vscode.template .vscode
Copy-Item -r bin.template bin

# $Env:CGO_ENABLED=0
# go get -u -v github.com/smallnest/gen