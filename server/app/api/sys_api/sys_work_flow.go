package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags workflow
// @Summary 注册工作流
// @Produce  application/json
// @Param data body sys_model.SysWorkflow true "注册工作流接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /workflow/createWorkFlow [post]
func CreateWorkFlow(c *gin.Context) {
	var wk sys_model.SysWorkflow
	_ = c.ShouldBindJSON(&wk)
	WKVerify := utils.Rules{
		"WorkflowNickName":    {utils.NotEmpty()},
		"WorkflowName":        {utils.NotEmpty()},
		"WorkflowDescription": {utils.NotEmpty()},
		"WorkflowStepInfo":    {utils.NotEmpty()},
	}
	WKVerifyErr := utils.Verify(wk, WKVerify)
	if WKVerifyErr != nil {
		respwrite.FailWithMessage(WKVerifyErr.Error(), c)
		return
	}
	err := sys_service.Create(wk)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		respwrite.OkWithMessage("获取成功", c)
	}
}
