package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags casbin
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "更改角色api权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/UpdateCasbin [post]
func UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		respwrite.FailWithMessage(AuthorityIdVerifyErr.Error(), c)
		return
	}
	err := sys_service.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("添加规则失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("添加规则成功", c)
	}
}

// @Tags casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func GetPolicyPathByAuthorityId(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		respwrite.FailWithMessage(AuthorityIdVerifyErr.Error(), c)
		return
	}
	paths := sys_service.GetPolicyPathByAuthorityId(cmr.AuthorityId)
	respwrite.OkWithData(response.PolicyPathResponse{Paths: paths}, c)
}

// @Tags casbin
// @Summary casb RBAC RESTFUL测试路由
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/CasbinTest [get]
func CasbinTest(c *gin.Context) {
	// 测试restful以及占位符代码  随意书写
	pathParam := c.Param("pathParam")
	query := c.Query("query")
	respwrite.OkDetailed(gin.H{"pathParam": pathParam, "query": query}, "获取规则成功", c)
}
