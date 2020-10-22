package cache

import (
	"daigou-admin/config"
	"daigou-admin/global"
	"daigou-admin/global/giface"
	"daigou-admin/utils/library"
	"daigou-admin/utils/zaplog"
)

type CacheManage struct {
	giface.CacheManageer
	sysRedis *library.Redis //管理后台redis
	serRedis *library.Redis //业务redis
}

func NewCacheManage() *CacheManage {
	return &CacheManage{}
}

func (r *CacheManage) Setup() {

	//admin缓存
	sysCache := global.G_Config.Redis[0]
	sysCache.Type = config.SysCache
	r.Connect(sysCache, 0)

	//业务缓存
	serConf := global.G_Config.Service[0]
	serCache := &config.Redis{
		URL:  serConf.Redis,
		Name: serConf.Name,
		Type: config.SerCache,
	}
	r.Connect(serCache, 0)

	global.G_CacheManage = r
	zaplog.Info("initCacheManagee success")
}

func (r *CacheManage) Connect(v *config.Redis, index int) (err error) {

	if v.Type == config.SerCache {
		r.serRedis, err = library.NewRedis(&library.RedisConf{
			Index: index,
			Redis: v,
		})
	} else if v.Type == config.SysCache {
		r.sysRedis, err = library.NewRedis(&library.RedisConf{
			Index: index,
			Redis: v,
		})
	}
	if err != nil {
		zaplog.Errorf("Setup err:%s", err)
		return
	}
	return
}

func (r *CacheManage) Close() {
	r.serRedis.Close()
	r.sysRedis.Close()
}

//获取业务数缓存
func (r *CacheManage) GetSerCache() *library.Redis {
	return r.serRedis
}

//获取管理后台缓存
func (r *CacheManage) GetSysCache() *library.Redis {
	return r.sysRedis
}

//设置服务当前连接redis
func (r *CacheManage) ChangeSerRedis(index int) (err error) {
	if r.serRedis.Conf.Index == index {
		return
	}
	//关闭旧连接
	err = r.serRedis.Close()
	if err != nil {
		return
	}

	serConf := global.G_Config.Service[index]
	serCache := &config.Redis{
		URL:  serConf.Redis,
		Name: serConf.Name,
		Type: config.SerCache,
	}
	r.Connect(serCache, index)
	return
}
