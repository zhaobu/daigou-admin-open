package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionaryDetail true "创建SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func CreateSysDictionaryDetail(c *gin.Context) {
	var sysDictionaryDetail sys_model.SysDictionaryDetail
	_ = c.ShouldBindJSON(&sysDictionaryDetail)
	err := sys_service.CreateSysDictionaryDetail(sysDictionaryDetail)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 删除SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionaryDetail true "删除SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func DeleteSysDictionaryDetail(c *gin.Context) {
	var sysDictionaryDetail sys_model.SysDictionaryDetail
	_ = c.ShouldBindJSON(&sysDictionaryDetail)
	err := sys_service.DeleteSysDictionaryDetail(sysDictionaryDetail)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func UpdateSysDictionaryDetail(c *gin.Context) {
	var sysDictionaryDetail sys_model.SysDictionaryDetail
	_ = c.ShouldBindJSON(&sysDictionaryDetail)
	err := sys_service.UpdateSysDictionaryDetail(&sysDictionaryDetail)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 用id查询SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionaryDetail true "用id查询SysDictionaryDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func FindSysDictionaryDetail(c *gin.Context) {
	var sysDictionaryDetail sys_model.SysDictionaryDetail
	_ = c.ShouldBindQuery(&sysDictionaryDetail)
	err, resysDictionaryDetail := sys_service.GetSysDictionaryDetail(sysDictionaryDetail.ID)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{"resysDictionaryDetail": resysDictionaryDetail}, c)
	}
}

// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictionaryDetailSearch true "分页获取SysDictionaryDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := sys_service.GetSysDictionaryDetailInfoList(pageInfo)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
