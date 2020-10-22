package dg_api

import (
	"daigou-admin/app/service/dg_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"daigou-admin/utils/zaplog"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

//用户管理

// @Tags UserInfo
// @Summary 批量删除UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userManage/deleteUserInfoByIds [delete]
func DeleteUserInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := dg_service.DeleteUserInfoByIds(IDS)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags UserInfo
// @Summary 根据条件分页查找UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FindUserInfoListReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.FindUserInfoListResp} "请求结果"
// @Router /userManage/findUserInfoList [get]
func FindUserInfoList(c *gin.Context) {
	var (
		req  request.FindUserInfoListReq
		resp *response.FindUserInfoListResp
		err  error
	)
	//获取参数
	if err = c.ShouldBindJSON(&req); err != nil {
		zaplog.Errorf("api FindUserInfoList err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("FindUserInfoList request:%s", utils.Dump(req))

	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api FindUserInfoList err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	//业务处理
	err, resp = dg_service.FindUserInfoList(&req)
	if err != nil {
		zaplog.Errorf("api FindUserInfoList err:%s", err)
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}

	// zaplog.Debugf("FindUserInfoList response:%s", utils.Dump(resp))
	respwrite.OkWithData(response.PageResult{
		List:     resp.Info,
		Total:    resp.TotalCount,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

// @Tags UserInfo
// @Summary 操作店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperateUserReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.OperateUserResp} "请求结果"
// @Router /userManage/operateUser [PUT]
func OperateUser(c *gin.Context) {
	var (
		req  request.OperateUserReq
		resp *response.OperateUserResp
		err  error
	)
	//获取参数
	if err = c.ShouldBindJSON(&req); err != nil {
		zaplog.Errorf("api OperateUser err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("OperateUser request:%s", utils.Dump(req))

	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api OperateUser err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	//业务处理
	err, resp, msg := dg_service.OperateUser(&req)
	if err != nil {
		zaplog.Errorf("api OperateUser err:%s", err)
		respwrite.FailWithMessage(fmt.Sprintf("用户%d操作失败，%v", req.UserID, err), c)
		return
	}

	zaplog.Debugf("OperateUser response:%s", utils.Dump(resp))
	respwrite.OkDetailed(resp, msg, c)
}
