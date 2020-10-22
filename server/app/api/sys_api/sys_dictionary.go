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

// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionary true "创建SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/createSysDictionary [post]
func CreateSysDictionary(c *gin.Context) {
	var sysDictionary sys_model.SysDictionary
	_ = c.ShouldBindJSON(&sysDictionary)
	err := sys_service.CreateSysDictionary(sysDictionary)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionary true "删除SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionary/deleteSysDictionary [delete]
func DeleteSysDictionary(c *gin.Context) {
	var sysDictionary sys_model.SysDictionary
	_ = c.ShouldBindJSON(&sysDictionary)
	err := sys_service.DeleteSysDictionary(sysDictionary)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionary true "更新SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionary/updateSysDictionary [put]
func UpdateSysDictionary(c *gin.Context) {
	var sysDictionary sys_model.SysDictionary
	_ = c.ShouldBindJSON(&sysDictionary)
	err := sys_service.UpdateSysDictionary(&sysDictionary)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysDictionary true "用id查询SysDictionary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionary/findSysDictionary [get]
func FindSysDictionary(c *gin.Context) {
	var sysDictionary sys_model.SysDictionary
	_ = c.ShouldBindQuery(&sysDictionary)
	err, resysDictionary := sys_service.GetSysDictionary(sysDictionary.Type, sysDictionary.ID)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{"resysDictionary": resysDictionary}, c)
	}
}

// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictionarySearch true "分页获取SysDictionary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionary/getSysDictionaryList [get]
func GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := sys_service.GetSysDictionaryInfoList(pageInfo)
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
