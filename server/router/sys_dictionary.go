package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSysDictionaryRouter(Router *gin.RouterGroup) {
	SysDictionaryRouter := Router.Group("sysDictionary").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		SysDictionaryRouter.POST("createSysDictionary", sys_api.CreateSysDictionary)   // 新建SysDictionary
		SysDictionaryRouter.DELETE("deleteSysDictionary", sys_api.DeleteSysDictionary) // 删除SysDictionary
		SysDictionaryRouter.PUT("updateSysDictionary", sys_api.UpdateSysDictionary)    // 更新SysDictionary
		SysDictionaryRouter.GET("findSysDictionary", sys_api.FindSysDictionary)        // 根据ID获取SysDictionary
		SysDictionaryRouter.GET("getSysDictionaryList", sys_api.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}
