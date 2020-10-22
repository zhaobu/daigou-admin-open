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
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func CreateApi(c *gin.Context) {
	var api sys_model.SysApi
	_ = c.ShouldBindJSON(&api)
	ApiVerify := utils.Rules{
		"Path":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"ApiGroup":    {utils.NotEmpty()},
		"Method":      {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(api, ApiVerify)
	if ApiVerifyErr != nil {
		respwrite.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := sys_service.CreateApi(api)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func DeleteApi(c *gin.Context) {
	var a sys_model.SysApi
	_ = c.ShouldBindJSON(&a)
	ApiVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	if ApiVerifyErr != nil {
		respwrite.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := sys_service.DeleteApi(a)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// 条件搜索后端看此api

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func GetApiList(c *gin.Context) {
	// 此结构体仅本方法使用
	var sp request.SearchApiParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		respwrite.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := sys_service.GetAPIInfoList(sp.SysApi, sp.PageInfo, sp.OrderKey, sp.Desc)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     sp.PageInfo.Page,
			PageSize: sp.PageInfo.PageSize,
		}, c)
	}
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func GetApiById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		respwrite.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, api := sys_service.GetApiById(idInfo.Id)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.SysAPIResponse{Api: api}, c)
	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updateApi [post]
func UpdateApi(c *gin.Context) {
	var api sys_model.SysApi
	_ = c.ShouldBindJSON(&api)
	ApiVerify := utils.Rules{
		"Path":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"ApiGroup":    {utils.NotEmpty()},
		"Method":      {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(api, ApiVerify)
	if ApiVerifyErr != nil {
		respwrite.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := sys_service.UpdateApi(api)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("修改数据成功", c)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func GetAllApis(c *gin.Context) {
	err, apis := sys_service.GetAllApis()
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.SysAPIListResponse{Apis: apis}, c)
	}
}
