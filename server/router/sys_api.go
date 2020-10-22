package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("createApi", sys_api.CreateApi)   // 创建Api
		ApiRouter.POST("deleteApi", sys_api.DeleteApi)   // 删除Api
		ApiRouter.POST("getApiList", sys_api.GetApiList) // 获取Api列表
		ApiRouter.POST("getApiById", sys_api.GetApiById) // 获取单条Api消息
		ApiRouter.POST("updateApi", sys_api.UpdateApi)   // 更新api
		ApiRouter.POST("getAllApis", sys_api.GetAllApis) // 获取所有api
	}
}
