package dg_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
)

func GetConnect() (err error, resp *response.GetConnectResp) {

	resp = &response.GetConnectResp{
		CurrentIndex:    global.G_CacheManage.GetSerCache().Conf.Index,
		RedisConfigList: make(map[int]*response.RedisConfig, len(global.G_Config.Service)),
		MysqlConfigList: make(map[int]*response.MysqlConfig, len(global.G_Config.Service)),
	}

	for index, v := range global.G_Config.Service {

		options, err := redis.ParseURL(v.Redis)
		if err != nil {
			err = fmt.Errorf("URL 解析错误:%w.", err)
			return err, nil
		}
		resp.RedisConfigList[index] = &response.RedisConfig{
			Name:     v.Name,
			Addr:     options.Addr,
			Password: options.Password,
			DB:       options.DB,
			Index:    index,
		}

		options1, err := mysql.ParseDSN(v.DB)
		if err != nil {
			err = fmt.Errorf("URL 解析错误:%w.", err)
			return err, nil
		}
		resp.MysqlConfigList[index] = &response.MysqlConfig{
			Name:     v.Name,
			UserName: options1.User,
			Password: options1.Passwd,
			Addr:     options1.Addr,
			DBName:   options1.DBName,
			Params:   options1.Params,
			Index:    index,
		}
	}

	return
}

func ChangeConnect(req *request.ChangeConnectReq) (err error, msg string) {

	if req.Index < 0 || req.Index >= len(global.G_Config.Service) {
		err = fmt.Errorf("%w", errors.New("所选择的配置不存在"))
		return
	}
	err = global.G_CacheManage.ChangeSerRedis(req.Index)
	if err != nil {
		err = fmt.Errorf("切换redis失败,%w", err)
		return
	}
	err = global.G_DBManage.ChangeSerDB(req.Index)
	if err != nil {
		err = fmt.Errorf("切换mysql失败,%w", err)
		return
	}
	msg = "切换连接成功"
	return
}

// func GetRedisList() (err error, resp *response.GetRedisListResp) {

// 	resp = &response.GetRedisListResp{
// 		CurrentIndex:    global.G_CacheManage.GetSerCache().Conf.Index,
// 		RedisConfigList: make(map[int]*response.RedisConfig, len(global.G_Config.Service)),
// 	}

// 	for index, v := range global.G_Config.Service {

// 		options, err := redis.ParseURL(v.Redis)
// 		if err != nil {
// 			err = fmt.Errorf("URL 解析错误:%w.", err)
// 			return err, nil
// 		}
// 		resp.RedisConfigList[index] = &response.RedisConfig{
// 			Name:     v.Name,
// 			Addr:     options.Addr,
// 			Password: options.Password,
// 			DB:       options.DB,
// 			Index:    index,
// 		}
// 	}

// 	return
// }

// func ChangeRedis(req *request.ChangeRedisReq) (err error, msg string) {

// 	if req.Index < 0 || req.Index >= len(global.G_Config.Service) {
// 		err = fmt.Errorf("%w", errors.New("所选择的配置不存在"))
// 		return
// 	}
// 	err = global.G_CacheManage.ChangeSerRedis(req.Index)
// 	if err != nil {
// 		err = fmt.Errorf("切换redis失败,%w", err)
// 		return
// 	}
// 	msg = "切换redis连接成功"
// 	return
// }

// func GetMysqlList() (err error, resp *response.GetMysqlListResp) {

// 	resp = &response.GetMysqlListResp{
// 		CurrentIndex:    global.G_DBManage.GetSerDB().Conf.Index,
// 		MysqlConfigList: make(map[int]*response.MysqlConfig, len(global.G_Config.Service)),
// 	}

// 	//当前连接的数据库
// 	for index, v := range global.G_Config.Service {

// 		options, err := mysql.ParseDSN(v.DB)
// 		if err != nil {
// 			err = fmt.Errorf("URL 解析错误:%w.", err)
// 			return err, nil
// 		}
// 		resp.MysqlConfigList[index] = &response.MysqlConfig{
// 			Name:     v.Name,
// 			UserName: options.User,
// 			Password: options.Passwd,
// 			Addr:     options.Addr,
// 			DBName:   options.DBName,
// 			Params:   options.Params,
// 			Index:    index,
// 		}
// 	}

// 	return
// }

// func ChangeMysql(req *request.ChangeMysqlReq) (err error, msg string) {

// 	if req.Index < 0 || req.Index >= len(global.G_Config.Service) {
// 		err = fmt.Errorf("%w", errors.New("所选择的配置不存在"))
// 		return
// 	}
// 	err = global.G_DBManage.ChangeSerDB(req.Index)
// 	if err != nil {
// 		err = fmt.Errorf("切换mysql失败,%w", err)
// 		return
// 	}
// 	msg = "切换mysql连接成功"
// 	return
// }
