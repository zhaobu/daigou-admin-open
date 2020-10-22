package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", sys_api.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", sys_api.GetPolicyPathByAuthorityId)
		CasbinRouter.GET("casbinTest/:pathParam", sys_api.CasbinTest)
	}
}
