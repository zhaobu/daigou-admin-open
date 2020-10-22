package router

import (
	"daigou-admin/app/api/sys_api"

	"github.com/gin-gonic/gin"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload")
	// .Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		FileUploadAndDownloadGroup.POST("/upload", sys_api.UploadFile)                                 // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", sys_api.GetFileList)                           // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", sys_api.DeleteFile)                             // 删除指定文件
		FileUploadAndDownloadGroup.POST("/breakpointContinue", sys_api.BreakpointContinue)             // 断点续传
		FileUploadAndDownloadGroup.GET("/findFile", sys_api.FindFile)                                  // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/breakpointContinueFinish", sys_api.BreakpointContinueFinish) // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/removeChunk", sys_api.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
