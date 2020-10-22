package sys_service

import (
	"daigou-admin/global"
	"daigou-admin/model/sys_model"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// 保存文件切片路径
func SaveChunk(uploader sys_model.ExaSimpleUploader) (err error) {
	return global.GetSysDB().Create(uploader).Error
}

// 检查文件是否已经上传过
func CheckFileMd5(md5 string) (err error, uploads []sys_model.ExaSimpleUploader, isDone bool) {
	err = global.GetSysDB().Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
	isDone = global.GetSysDB().First(&sys_model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).RecordNotFound()
	return err, uploads, !isDone
}

// 合并文件
func MergeFileMd5(md5 string, fileName string) (err error) {
	finishDir := "./finish/"
	dir := "./chunk/" + md5
	//如果文件上传成功 不做后续操作 通知成功即可
	notFinish := global.GetSysDB().First(&sys_model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).RecordNotFound()
	if !notFinish {
		return nil
	}

	//打开切片文件夹
	rd, err := ioutil.ReadDir(dir)
	_ = os.MkdirAll(finishDir, os.ModePerm)
	//创建目标文件
	fd, _ := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(dir + "/" + fileName + strconv.Itoa(k+1))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
		}
	}
	//关闭文件
	defer fd.Close()

	if err != nil {
		return err
	}
	//创建事务
	tx := global.GetSysDB().Begin()
	//删除切片信息
	err = tx.Delete(&sys_model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, false).Error
	// 添加文件信息
	err = tx.Create(&sys_model.ExaSimpleUploader{
		Identifier: md5,
		IsDone:     true,
		FilePath:   finishDir + fileName,
		Filename:   fileName,
	}).Error
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	tx.Commit()
	//清除切片
	err = os.RemoveAll(dir)
	return
}
