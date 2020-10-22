package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		UserRouter.POST("changePassword", sys_api.ChangePassword)     // 修改密码
		UserRouter.POST("uploadHeaderImg", sys_api.UploadHeaderImg)   // 上传头像
		UserRouter.POST("getUserList", sys_api.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", sys_api.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", sys_api.DeleteUser)           // 删除用户
	}
}
