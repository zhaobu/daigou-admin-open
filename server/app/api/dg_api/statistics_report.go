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

//统计报表

// @Tags statics_report
// @Summary 新增用户
// @Description 查询新增用户
// @Produce  application/json
// @Accept  application/json
// @Param data body request.NewUserReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.NewUserResp} "请求结果"
// @Router /daigou/staticsReport/newUser [GET]
func NewUser(c *gin.Context) {
	var (
		req  request.NewUserReq
		resp []*response.NewUserResp
		err  error
	)
	//获取参数
	if err = c.ShouldBindQuery(&req); err != nil {
		zaplog.Errorf("api NewUser err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api NewUser err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	zaplog.Debugf("NewUser request:%s", utils.Dump(req))
	//业务处理
	err, resp = dg_service.NewUser(&req)
	if err != nil {
		zaplog.Errorf("api NewUser err:%s", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}

	zaplog.Debugf("NewUser response:%s", utils.Dump(resp))
	respwrite.OkWithData(resp, c)
}

// @Tags statics_report
// @Summary 活跃用户
// @Description 查询活跃用户
// @Produce  application/json
// @Accept  application/json
// @Param data body request.ActiveUserReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.ActiveUserResp} "请求结果"
// @Router /daigou/staticsReport/activeUser [GET]
func ActiveUser(c *gin.Context) {
	var (
		req  request.ActiveUserReq
		resp []*response.ActiveUserResp
		err  error
	)
	//获取参数
	if err = c.ShouldBindQuery(&req); err != nil {
		zaplog.Errorf("api ActiveUser err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api ActiveUser err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	zaplog.Debugf("ActiveUser request:%s", utils.Dump(req))
	//业务处理
	err, resp = dg_service.ActiveUser(&req)
	if err != nil {
		zaplog.Errorf("api ActiveUser err:%s", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}

	zaplog.Debugf("ActiveUser response:%s", utils.Dump(resp))
	respwrite.OkWithData(resp, c)
}
