package initialize

import (
	"daigou-admin/global"
	"daigou-admin/pkg/authcasbin"
	"daigou-admin/pkg/cache"
	"daigou-admin/pkg/database"
	"daigou-admin/pkg/logger"
)

func Init() {
	//初始化配置文件
	initConfig()

	//初始化日志
	initlogger()

	//初始化参数检查
	initValidator()

	//初始化db
	initDBMange()

	//初始化redis
	initCacheManage()

	//初始化casbin
	initCasbin()
}

func initDBMange() {
	dbman := database.NewDBMange()
	dbman.Setup()
	global.G_DBManage = dbman
}

func initCacheManage() {
	cacheMan := cache.NewCacheManage()
	cacheMan.Setup()
	global.G_CacheManage = cacheMan
}

func initlogger() {
	logger := logger.NewLogger()
	logger.Setup()
	global.G_Logger = logger
}

func initCasbin() {
	auth := authcasbin.NewCasbin()
	auth.Setup()
	global.G_Casbin = auth
}
