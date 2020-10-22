package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("getSystemConfig", sys_api.GetSystemConfig) // 获取配置文件内容
		UserRouter.POST("setSystemConfig", sys_api.SetSystemConfig) // 设置配置文件内容
	}
}
