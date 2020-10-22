package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSimpleUploaderRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("simpleUploader").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("upload", sys_api.SimpleUploaderUpload) // 上传功能
		ApiRouter.GET("checkFileMd5", sys_api.CheckFileMd5)    // 文件完整度验证
		ApiRouter.GET("mergeFileMd5", sys_api.MergeFileMd5)    // 合并文件
	}
}
