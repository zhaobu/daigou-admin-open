package library

import (
	"context"
	"daigou-admin/config"
	"daigou-admin/utils/zaplog"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/os/gtimer"
)

type Redis struct {
	*redis.Client
	pingTimer *gtimer.Entry
	Conf      *RedisConf
}

type RedisConf struct {
	*config.Redis
	Index int
}

func (r Redis) Close() error {
	r.pingTimer.Close()
	return r.Client.Close()
}

func NewRedis(conf *RedisConf) (res *Redis, err error) {
	redisCfg, err := redis.ParseURL(conf.URL)
	if err != nil {
		return nil, err
	}
	res = &Redis{
		Conf: conf,
	}
	retryTimes := 60
	for i := 0; i < retryTimes; i++ {
		zaplog.Infof("try to connect redis %s the %d time", conf.URL, i)
		res.Client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password, // no password set
			DB:       redisCfg.DB,       // use default DB
		})
		_, err := res.Client.Ping(context.Background()).Result() //测试是否连接成功
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	res.pingTimer = gtimer.AddSingleton(time.Minute, func() {
		_, err := res.Client.Ping(context.Background()).Result()
		if err != nil {
			zaplog.Errorf("pingRedis err=%s", err)
			return
		}
	})

	zaplog.Infof("redis %s connect success:%s ", conf.Name, conf.URL)
	return
}
