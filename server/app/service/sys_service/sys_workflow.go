package sys_service

import (
	"daigou-admin/global"
	"daigou-admin/model/sys_model"
)

// @title    Create
// @description   create a workflow, 创建工作流
// @auth                     （2020/04/05  20:22）
// @param     wk              sys_model.SysWorkflow
// @return                    error

func Create(wk sys_model.SysWorkflow) error {
	err := global.GetSysDB().Create(&wk).Error
	return err
}
