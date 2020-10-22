package giface

import "daigou-admin/utils/library"

//缓存
type CacheManageer interface {
	Setup()
	Close()
	//获取管理后台缓存
	GetSysCache() *library.Redis
	//获取业务数缓存
	GetSerCache() *library.Redis
	//设置服务当前连接redis
	ChangeSerRedis(index int) error
}
