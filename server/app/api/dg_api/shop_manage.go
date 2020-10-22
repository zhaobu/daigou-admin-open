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

//店铺管理

// @Tags ShopInfo
// @Summary 批量删除ShopInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ShopInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopManage/deleteShopInfoByIds [delete]
func DeleteShopInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := dg_service.DeleteShopInfoByIds(IDS)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags ShopInfo
// @Summary 根据条件分页查找ShopInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FindShopInfoListReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.FindShopInfoListResp} "请求结果"
// @Router /shopManage/findShopInfoList [get]
func FindShopInfoList(c *gin.Context) {
	var (
		req  request.FindShopInfoListReq
		resp *response.FindShopInfoListResp
		err  error
	)
	//获取参数
	if err = c.ShouldBindJSON(&req); err != nil {
		zaplog.Errorf("api FindShopInfoList err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("FindShopInfoList request:%s", utils.Dump(req))

	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api FindShopInfoList err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	//业务处理
	err, resp = dg_service.FindShopInfoList(&req)
	if err != nil {
		zaplog.Errorf("api FindShopInfoList err:%s", err)
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}

	zaplog.Debugf("FindShopInfoList response:%s", utils.Dump(resp))
	respwrite.OkWithData(response.PageResult{
		List:     resp.Info,
		Total:    resp.TotalCount,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

// @Tags ShopInfo
// @Summary 操作店铺
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperateShopReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.OperateShopResp} "请求结果"
// @Router /shopManage/operateShop [PUT]
func OperateShop(c *gin.Context) {
	var (
		req  request.OperateShopReq
		resp *response.OperateShopResp
		err  error
	)
	//获取参数
	if err = c.ShouldBindJSON(&req); err != nil {
		zaplog.Errorf("api OperateShop err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("OperateShop request:%s", utils.Dump(req))

	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api OperateShop err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	//业务处理
	err, resp, msg := dg_service.OperateShop(&req)
	if err != nil {
		zaplog.Errorf("api OperateShop err:%s", err)
		respwrite.FailWithMessage(fmt.Sprintf("店铺%d操作失败，%v", req.ShopID, err), c)
		return
	}

	zaplog.Debugf("OperateShop response:%s", utils.Dump(resp))
	respwrite.OkDetailed(resp, msg, c)
}
