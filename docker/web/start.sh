#!/bin/sh

npm config set prefix "/usr/local/node_global"  
npm config set cache "/usr/local/node_cache"
# 配置镜像站
npm config set registry=http://registry.npm.taobao.org
# 安装必要包
cd code
yarn install 
npm run build


# node image 官方启动脚本
# ./docker-entrypoint.sh