package router

import (
	"daigou-admin/app/api/sys_api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", sys_api.Register)
		BaseRouter.POST("login", sys_api.Login)
		BaseRouter.POST("captcha", sys_api.Captcha)
	}
	return BaseRouter
}
