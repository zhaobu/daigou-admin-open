package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("sysOperationRecord").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler())
	{
		SysOperationRecordRouter.POST("createSysOperationRecord", sys_api.CreateSysOperationRecord)             // 新建SysOperationRecord
		SysOperationRecordRouter.DELETE("deleteSysOperationRecord", sys_api.DeleteSysOperationRecord)           // 删除SysOperationRecord
		SysOperationRecordRouter.DELETE("deleteSysOperationRecordByIds", sys_api.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		SysOperationRecordRouter.PUT("updateSysOperationRecord", sys_api.UpdateSysOperationRecord)              // 更新SysOperationRecord
		SysOperationRecordRouter.GET("findSysOperationRecord", sys_api.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		SysOperationRecordRouter.GET("getSysOperationRecordList", sys_api.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
