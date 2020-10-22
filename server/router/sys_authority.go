package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		AuthorityRouter.POST("createAuthority", sys_api.CreateAuthority)   // 创建角色
		AuthorityRouter.POST("deleteAuthority", sys_api.DeleteAuthority)   // 删除角色
		AuthorityRouter.PUT("updateAuthority", sys_api.UpdateAuthority)    // 更新角色
		AuthorityRouter.POST("copyAuthority", sys_api.CopyAuthority)       // 更新角色
		AuthorityRouter.POST("getAuthorityList", sys_api.GetAuthorityList) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", sys_api.SetDataAuthority) // 设置角色资源权限
	}
}
