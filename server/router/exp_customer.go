package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("customer").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		ApiRouter.POST("customer", sys_api.CreateExaCustomer)     // 创建客户
		ApiRouter.PUT("customer", sys_api.UpdateExaCustomer)      // 更新客户
		ApiRouter.DELETE("customer", sys_api.DeleteExaCustomer)   // 删除客户
		ApiRouter.GET("customer", sys_api.GetExaCustomer)         // 获取单一客户信息
		ApiRouter.GET("customerList", sys_api.GetExaCustomerList) // 获取客户列表
	}
}
