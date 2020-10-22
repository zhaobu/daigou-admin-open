package dg_api

import (
	"daigou-admin/app/service/dg_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"daigou-admin/utils/zaplog"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

//服务器配置

// @Tags 服务器配置
// @Summary 获取连接配置列表
// @Description 获取连接配置列表
// @Produce  application/json
// @Accept  application/json
// @Param data body request.GetConnectReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.GetConnectResp} "请求结果"
// @Router /daigou/serverConfig/getConnectList [GET]
func GetConnect(c *gin.Context) {
	var (
		resp *response.GetConnectResp
		err  error
	)

	//业务处理
	err, resp = dg_service.GetConnect()
	if err != nil {
		zaplog.Errorf("api GetConnect err:%s", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}

	zaplog.Debugf("GetConnect response:%s", utils.Dump(resp))
	respwrite.OkWithData(resp, c)
}

// @Tags 服务器配置
// @Summary 修改连接
// @Description 修改连接
// @Produce  application/json
// @Accept  application/json
// @Param data body request.ChangeConnectReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{} "请求结果"
// @Router /daigou/serverConfig/changeConnect [POST]
func ChangeConnect(c *gin.Context) {
	var (
		req request.ChangeConnectReq
		msg string
		err error
	)
	//获取参数
	if err = c.ShouldBindJSON(&req); err != nil {
		zaplog.Errorf("api ChangeConnect err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api ChangeConnect err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	zaplog.Debugf("ChangeConnect request:%s", utils.Dump(req))
	//业务处理
	err, msg = dg_service.ChangeConnect(&req)
	if err != nil {
		zaplog.Errorf("api ChangeConnect err:%s", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}

	zaplog.Debugf("ChangeConnect response:%s", msg)
	respwrite.OkWithMessage(msg, c)
}

// // @Tags 服务器配置
// // @Summary 获取redis配置列表
// // @Description 获取redis配置列表
// // @Produce  application/json
// // @Accept  application/json
// // @Param data body request.GetRedisListReq true "请求参数"
// // @Success 200 {object} respwrite.JsonResponse{data=response.GetRedisListResp} "请求结果"
// // @Router /daigou/serverConfig/getRedisList [GET]
// func GetRedisList(c *gin.Context) {
// 	var (
// 		resp *response.GetRedisListResp
// 		err  error
// 	)

// 	//业务处理
// 	err, resp = dg_service.GetRedisList()
// 	if err != nil {
// 		zaplog.Errorf("api GetRedisList err:%s", err)
// 		respwrite.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	zaplog.Debugf("GetRedisList response:%s", utils.Dump(resp))
// 	respwrite.OkWithData(resp, c)
// }

// // @Tags 服务器配置
// // @Summary 修改redis连接
// // @Description 修改redis连接
// // @Produce  application/json
// // @Accept  application/json
// // @Param data body request.ChangeRedisReq true "请求参数"
// // @Success 200 {object} respwrite.JsonResponse{} "请求结果"
// // @Router /daigou/serverConfig/changeRedis [POST]
// func ChangeRedis(c *gin.Context) {
// 	var (
// 		req request.ChangeRedisReq
// 		msg string
// 		err error
// 	)
// 	//获取参数
// 	if err = c.ShouldBindJSON(&req); err != nil {
// 		zaplog.Errorf("api ChangeRedis err:%w", err)
// 		respwrite.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	//检查参数
// 	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
// 		if v, ok := err.(*gvalid.Error); ok {
// 			zaplog.Errorf("api ChangeRedis err:%s", v.FirstString())
// 			respwrite.FailWithMessage(v.FirstString(), c)
// 		}
// 		return
// 	}

// 	zaplog.Debugf("ChangeRedis request:%s", utils.Dump(req))
// 	//业务处理
// 	err, msg = dg_service.ChangeRedis(&req)
// 	if err != nil {
// 		zaplog.Errorf("api ChangeRedis err:%s", err)
// 		respwrite.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	zaplog.Debugf("ChangeRedis response:%s", msg)
// 	respwrite.OkWithMessage(msg, c)
// }

// // @Tags 服务器配置
// // @Summary 获取mysql配置列表
// // @Description 获取mysql配置列表
// // @Produce  application/json
// // @Accept  application/json
// // @Param data body request.GetMysqlListReq true "请求参数"
// // @Success 200 {object} respwrite.JsonResponse{data=response.GetMysqlListResp} "请求结果"
// // @Router /daigou/serverConfig/getMysqlList [GET]
// func GetMysqlList(c *gin.Context) {
// 	var (
// 		resp *response.GetMysqlListResp
// 		err  error
// 	)

// 	//业务处理
// 	err, resp = dg_service.GetMysqlList()
// 	if err != nil {
// 		zaplog.Errorf("api GetMysqlList err:%s", err)
// 		respwrite.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	zaplog.Debugf("GetMysqlList response:%s", utils.Dump(resp))
// 	respwrite.OkWithData(resp, c)
// }

// // @Tags 服务器配置
// // @Summary 修改mysql连接
// // @Description 修改mysql连接
// // @Produce  application/json
// // @Accept  application/json
// // @Param data body request.ChangeMysqlReq true "请求参数"
// // @Success 200 {object} respwrite.JsonResponse{} "请求结果"
// // @Router /daigou/serverConfig/changeMysql [POST]
// func ChangeMysql(c *gin.Context) {
// 	var (
// 		req request.ChangeMysqlReq
// 		msg string
// 		err error
// 	)
// 	//获取参数
// 	if err = c.ShouldBindJSON(&req); err != nil {
// 		zaplog.Errorf("api ChangeMysql err:%w", err)
// 		respwrite.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	//检查参数
// 	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
// 		if v, ok := err.(*gvalid.Error); ok {
// 			zaplog.Errorf("api ChangeMysql err:%s", v.FirstString())
// 			respwrite.FailWithMessage(v.FirstString(), c)
// 		}
// 		return
// 	}

// 	zaplog.Debugf("ChangeMysql request:%s", utils.Dump(req))
// 	//业务处理
// 	err, msg = dg_service.ChangeMysql(&req)
// 	if err != nil {
// 		zaplog.Errorf("api ChangeMysql err:%s", err)
// 		respwrite.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	zaplog.Debugf("ChangeMysql response:%s", msg)
// 	respwrite.OkWithMessage(msg, c)
// }
