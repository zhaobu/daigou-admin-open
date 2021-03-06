package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
func CreateExaCustomer(c *gin.Context) {
	var cu sys_model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	CustomerVerify := utils.Rules{
		"CustomerName":      {utils.NotEmpty()},
		"CustomerPhoneData": {utils.NotEmpty()},
	}
	CustomerVerifyErr := utils.Verify(cu, CustomerVerify)
	if CustomerVerifyErr != nil {
		respwrite.FailWithMessage(CustomerVerifyErr.Error(), c)
		return
	}
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	cu.SysUserID = waitUse.ID
	cu.SysUserAuthorityID = waitUse.AuthorityId
	err := sys_service.CreateExaCustomer(cu)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		respwrite.OkWithMessage("创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
func DeleteExaCustomer(c *gin.Context) {
	var cu sys_model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	CustomerVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	CustomerVerifyErr := utils.Verify(cu.Model, CustomerVerify)
	if CustomerVerifyErr != nil {
		respwrite.FailWithMessage(CustomerVerifyErr.Error(), c)
		return
	}
	err := sys_service.DeleteExaCustomer(cu)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
func UpdateExaCustomer(c *gin.Context) {
	var cu sys_model.ExaCustomer
	_ = c.ShouldBindJSON(&cu)
	IdCustomerVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	IdCustomerVerifyErr := utils.Verify(cu.Model, IdCustomerVerify)
	if IdCustomerVerifyErr != nil {
		respwrite.FailWithMessage(IdCustomerVerifyErr.Error(), c)
		return
	}
	CustomerVerify := utils.Rules{
		"CustomerName":      {utils.NotEmpty()},
		"CustomerPhoneData": {utils.NotEmpty()},
	}
	CustomerVerifyErr := utils.Verify(cu, CustomerVerify)
	if CustomerVerifyErr != nil {
		respwrite.FailWithMessage(CustomerVerifyErr.Error(), c)
		return
	}
	err := sys_service.UpdateExaCustomer(&cu)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("更新失败：%v", err), c)
	} else {
		respwrite.OkWithMessage("更新成功", c)
	}
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func GetExaCustomer(c *gin.Context) {
	var cu sys_model.ExaCustomer
	_ = c.ShouldBindQuery(&cu)
	IdCustomerVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	IdCustomerVerifyErr := utils.Verify(cu.Model, IdCustomerVerify)
	if IdCustomerVerifyErr != nil {
		respwrite.FailWithMessage(IdCustomerVerifyErr.Error(), c)
		return
	}
	err, customer := sys_service.GetExaCustomer(cu.ID)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		respwrite.OkWithData(response.ExaCustomerResponse{Customer: customer}, c)
	}
}

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func GetExaCustomerList(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		respwrite.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, customerList, total := sys_service.GetCustomerInfoList(waitUse.AuthorityId, pageInfo)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		respwrite.OkWithData(response.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
