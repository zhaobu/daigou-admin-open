# node.js 环境安装

1. 去官网下载 node.js 安装包,比如安装到 D:\develope\nodejs
2. 在安装目录 D:\develope\nodejs 下新建 node_global(全局包存放目录) 和 node_cache(缓存目录)
3. 打开命令行工具，执行以下两句操作： `npm config set prefix "D:\develope\nodejs\node_global"` `npm config set cache "D:\develope\nodejs\node_cache"`
4. 打开系统属性-高级-环境变量，在系统变量中新建 变量名：NODE_PATH, 变量值： `D:\develope\nodejs\node_global\node_modules`
5. 编辑用户变量的 path，将默认的 C 盘下 `APPData/Roaming/npm` 修改为 `D:\develope\nodejs\node_global`
6. 使用 `npm config set registry https://registry.npm.taobao.org` 设置 NPM 源为淘宝源, (如果想还原 npm 官方仓库, 执行 `npm config set registry https://registry.npmjs.org/` )

# npm 设置代理

- 设置代理:

```shell
npm config set proxy=http://127.0.0.1:10809
npm config set https-proxy http://127.0.0.1:10809
```

- 取消代理:

```shell
npm config delete proxy
npm config delete https-proxy
```

# 安装 nrm 镜像管理工具

## install

```shell
npm install -g nrm
```

## example

```shell
nrm ls
npm -----  https://registry.npmjs.org/
yarn ----- https://registry.yarnpkg.com
cnpm ----  http://r.cnpmjs.org/
taobao --  https://registry.npm.taobao.org/
nj ------  https://registry.nodejitsu.com/
skimdb -- https://skimdb.npmjs.com/registry

nrm use cnpm  //switch registry to cnpm
```

# 依赖更新

参考[npm 依赖包的安装、更新、删除](https://www.jianshu.com/p/9b9166f7559c)

# npm 安装时报错

当使用 npm 进行安装的时候，如果出现：cb() nerver called 的报错，

1. 先使用：npm cache verify
2. 在使用：npm cache clean
3. 可能需要：npm cache clean -—force

可能需要 git config --local http.sslverify "false"

重新进行安装需要的包即可

# yarn 命令

1. 查看一下当前源

```shell
yarn config get registry
```

2. 切换为淘宝源和自带

```shell
yarn config set registry https://registry.npm.taobao.org
yarn config set registry https://registry.yarnpkg.com
```

3. yarn 对比 npm 常用命令

| 作用           | npm                                | Yarn                  |
| -------------- | ---------------------------------- | --------------------- |
| 安装           | npm install(i)                     | yarn                  |
| 卸载           | npm uninstall(un)                  | yarn remove           |
| 全局安装       | npm install xxx –-global(-g)       | yarn global add xxx   |
| 安装包         | npm install xxx –save(-S)          | yarn add xxx          |
| 开发模式安装包 | npm install xxx –save-dev(-D)      | yarn add xxx –dev(-D) |
| 更新           | npm update –save                   | yarn upgrade          |
| 全局更新       | npm update –global                 | yarn global upgrade   |
| 卸载           | npm uninstall \[–save/–save-dev\]  | yarn remove xx        |
| 清除缓存       | npm cache clean                    | yarn cache clean      |
| 重装           | rm -rf node_modules && npm install | yarn upgrade          |
