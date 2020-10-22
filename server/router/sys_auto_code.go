package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		AutoCodeRouter.POST("createTemp", sys_api.CreateTemp)    // 创建自动化代码
		AutoCodeRouter.GET("getTables", sys_api.GetTables)       // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", sys_api.GetDB)               // 获取数据库
		AutoCodeRouter.GET("getColume", sys_api.GetColume)       // 获取指定表所有字段信息
		AutoCodeRouter.GET("getDBConnect", sys_api.GetDBConnect) // 获取当前业务数据库和后台数据库连接
	}
}
