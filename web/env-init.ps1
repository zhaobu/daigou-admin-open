# 复制模版文件
if (Test-Path .vscode) { (Remove-Item -r .vscode) }

Copy-Item -r .vscode.template .vscode