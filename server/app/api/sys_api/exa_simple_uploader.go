package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传插件版示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /simpleUploader/upload [post]
func SimpleUploaderUpload(c *gin.Context) {
	var chunk sys_model.ExaSimpleUploader
	_, header, err := c.Request.FormFile("file")
	chunk.Filename = c.PostForm("filename")
	chunk.ChunkNumber = c.PostForm("chunkNumber")
	chunk.CurrentChunkSize = c.PostForm("currentChunkSize")
	chunk.Identifier = c.PostForm("identifier")
	chunk.TotalSize = c.PostForm("totalSize")
	chunk.TotalChunks = c.PostForm("totalChunks")
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		utils.CreateDir(chunkDir)
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err = c.SaveUploadedFile(header, chunkPath)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("切片创建失败，%v", err), c)
		return
	}
	chunk.CurrentChunkPath = chunkPath
	err = sys_service.SaveChunk(chunk)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("切片创建失败，%v", err), c)
		return
	} else {
		respwrite.Ok(c)
	}
}

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth

// @Produce  application/json
// @Param params md5 get "测试文件是否已经存在和判断已经上传过的切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /simpleUploader/checkFileMd5 [get]
func CheckFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	err, chunks, isDone := sys_service.CheckFileMd5(md5)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("md5读取失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{
			"chunks": chunks,
			"isDone": isDone,
		}, c)
	}
}

// @Tags SimpleUploader
// @Summary 合并文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param params md5 get "合并文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"合并成功"}"
// @Router /simpleUploader/mergeFileMd5 [get]
func MergeFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	fileName := c.Query("fileName")
	err := sys_service.MergeFileMd5(md5, fileName)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("md5读取失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{}, c)
	}
}
