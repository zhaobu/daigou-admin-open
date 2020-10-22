package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	SysDictionaryDetailRouter := Router.Group("sysDictionaryDetail").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		SysDictionaryDetailRouter.POST("createSysDictionaryDetail", sys_api.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		SysDictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", sys_api.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		SysDictionaryDetailRouter.PUT("updateSysDictionaryDetail", sys_api.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
		SysDictionaryDetailRouter.GET("findSysDictionaryDetail", sys_api.FindSysDictionaryDetail)        // 根据ID获取SysDictionaryDetail
		SysDictionaryDetailRouter.GET("getSysDictionaryDetailList", sys_api.GetSysDictionaryDetailList)  // 获取SysDictionaryDetail列表
	}
}
