package sys_service

import (
	"context"
	"daigou-admin/global"
	"daigou-admin/model/sys_model"
)

// @title    JsonInBlacklist
// @description   create jwt blacklist
// @param     jwtList         sys_model.JwtBlacklist
// @auth                     （2020/04/05  20:22）
// @return    err             error

func JsonInBlacklist(jwtList sys_model.JwtBlacklist) (err error) {
	err = global.GetSysDB().Create(&jwtList).Error
	return
}

// @title    IsBlacklist
// @description   check if the Jwt is in the blacklist or not, 判断JWT是否在黑名单内部
// @auth                     （2020/04/05  20:22）
// @param     jwt             string
// @param     jwtList         sys_model.JwtBlacklist
// @return    err             error

func IsBlacklist(jwt string, jwtList sys_model.JwtBlacklist) bool {
	isNotFound := global.GetSysDB().Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}

// @title    GetRedisJWT
// @description   Get user info in redis
// @auth                     （2020/04/05  20:22）
// @param     userName        string
// @return    err             error
// @return    redisJWT        string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.G_CacheManage.GetSysCache().Get(context.Background(), userName).Result()
	return err, redisJWT
}

// @title    SetRedisJWT
// @description   set jwt into the Redis
// @auth                     （2020/04/05  20:22）
// @param     jwtList         sys_model.JwtBlacklist
// @param     userName        string
// @return    err             error

func SetRedisJWT(jwtList sys_model.JwtBlacklist, userName string) (err error) {
	err = global.G_CacheManage.GetSysCache().Set(context.Background(), userName, jwtList.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
