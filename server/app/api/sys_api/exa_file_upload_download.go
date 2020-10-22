package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		// 文件上传后拿到文件路径
		err, filePath, key := utils.Upload(header)
		if err != nil {
			respwrite.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			var file sys_model.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = header.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			if noSave == "0" {
				err = sys_service.Upload(file)
			}
			if err != nil {
				respwrite.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				respwrite.OkDetailed(response.ExaFileResponse{File: file}, "上传成功", c)
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sys_model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func DeleteFile(c *gin.Context) {
	var file sys_model.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	err, f := sys_service.FindFile(file.ID)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		err = utils.DeleteFile(f.Key)
		if err != nil {
			respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)

		} else {
			err = sys_service.DeleteFile(f)
			if err != nil {
				respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
			} else {
				respwrite.OkWithMessage("删除成功", c)
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取文件户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := sys_service.GetFileRecordInfoList(pageInfo)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
