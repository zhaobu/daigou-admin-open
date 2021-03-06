package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/response"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags system
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	err, config := sys_service.GetSystemConfig()
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.SysConfigResponse{Config: config}, c)
	}
}

// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sys_model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/setSystemConfig [post]
func SetSystemConfig(c *gin.Context) {
	var sys sys_model.System
	_ = c.ShouldBindJSON(&sys)
	err := sys_service.SetSystemConfig(sys)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("设置失败，%v", err), c)
	} else {
		respwrite.OkWithData("设置成功", c)
	}
}

// 本方法开发中 开发者windows系统 缺少linux系统所需的包 因此搁置
// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sys_model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/ReloadSystem [post]
func ReloadSystem(c *gin.Context) {
	var sys sys_model.System
	_ = c.ShouldBindJSON(&sys)
	err := sys_service.SetSystemConfig(sys)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("设置失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("设置成功", c)
	}
}
