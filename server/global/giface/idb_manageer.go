package giface

import (
	"daigou-admin/utils/library"
)

// 数据库
type DBManageer interface {
	//初始化
	Setup()
	//初始化数据库
	AutoMigrate() error
	//获取管理后台数据库
	GetSysDB() *library.Mysql
	//获取业务数据库
	GetSerDB() *library.Mysql
	//设置服务当前连接mysql
	ChangeSerDB(index int) error
	//关闭数据库
	Close()
}
