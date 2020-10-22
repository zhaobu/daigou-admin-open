package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	modelJwt := sys_model.JwtBlacklist{
		Jwt: token,
	}
	err := sys_service.JsonInBlacklist(modelJwt)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("jwt作废失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("jwt作废成功", c)
	}
}
