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

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord sys_model.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	err := sys_service.CreateSysOperationRecord(sysOperationRecord)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("创建成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysOperationRecord true "删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord sys_model.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	err := sys_service.DeleteSysOperationRecord(sysOperationRecord)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := sys_service.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 更新SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysOperationRecord true "更新SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysOperationRecord/updateSysOperationRecord [put]
func UpdateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord sys_model.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	err := sys_service.UpdateSysOperationRecord(&sysOperationRecord)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("更新成功", c)
	}
}

// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysOperationRecord true "用id查询SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord sys_model.SysOperationRecord
	_ = c.ShouldBindQuery(&sysOperationRecord)
	err, resysOperationRecord := sys_service.GetSysOperationRecord(sysOperationRecord.ID)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{"resysOperationRecord": resysOperationRecord}, c)
	}
}

// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysOperationRecordSearch true "分页获取SysOperationRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func GetSysOperationRecordList(c *gin.Context) {
	var pageInfo request.SysOperationRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := sys_service.GetSysOperationRecordInfoList(pageInfo)
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
