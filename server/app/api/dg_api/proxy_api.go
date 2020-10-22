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

//代理访问

// @Tags 代理post访问
// @Summary admin代理访问server
// @Description admin使用post请求代理访问server
// @Produce  application/json
// @Accept  application/json
// @Param data body request.ProxyPostReq true "请求参数"
// @Success 200 {object} respwrite.JsonResponse{data=response.ProxyPostResp} "请求结果"
// @Router /daigou/proxyApi/proxyPost [POST]
func ProxyPost(c *gin.Context) {
	var (
		req  request.ProxyPostReq
		resp *response.ProxyPostResp
		err  error
	)

	//获取参数
	if err = c.ShouldBindJSON(&req); err != nil {
		zaplog.Errorf("api ProxyPost err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("ProxyPost request:%s", utils.Dump(req))

	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api ProxyPost err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	//业务处理
	err, resp, msg := dg_service.ProxyPost(&req)
	if err != nil {
		zaplog.Errorf("api ProxyPost err:%s", err)
		respwrite.FailWithMessage(fmt.Sprintf("代理请求失败,%v", err), c)
		return
	}

	zaplog.Debugf("ProxyPost response:%s", utils.Dump(resp))
	respwrite.OkDetailed(resp, msg, c)
}
