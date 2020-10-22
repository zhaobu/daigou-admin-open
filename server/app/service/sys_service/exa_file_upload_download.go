package sys_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/model/sys_model"
)

// @title    Upload
// @description   创建文件上传记录
// @param     file            sys_model.ExaFileUploadAndDownload
// @auth                     （2020/04/05  20:22）
// @return                    error

func Upload(file sys_model.ExaFileUploadAndDownload) error {
	err := global.GetSysDB().Create(&file).Error
	return err
}

// @title    FindFile
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error

func FindFile(id uint) (error, sys_model.ExaFileUploadAndDownload) {
	var file sys_model.ExaFileUploadAndDownload
	err := global.GetSysDB().Where("id = ?", id).First(&file).Error
	return err, file
}

// @title    DeleteFile
// @description   删除文件记录
// @auth                     （2020/04/05  20:22）
// @param     file            sys_model.ExaFileUploadAndDownload
// @return                    error

func DeleteFile(file sys_model.ExaFileUploadAndDownload) error {
	err := global.GetSysDB().Where("id = ?", file.ID).Unscoped().Delete(file).Error
	return err
}

// @title    GetFileRecordInfoList
// @description   分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return    err             error
// @return    list            error
// @return    total           error

func GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GetSysDB()
	var fileLists []sys_model.ExaFileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}
