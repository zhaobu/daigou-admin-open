package sys_service

import (
	"daigou-admin/config"
	"daigou-admin/global"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
)

// @title    GetSystemConfig
// @description   读取配置文件
// @auth                     （2020/04/05  20:22）
// @return    err             error
// @return    conf            Server

func GetSystemConfig() (err error, conf *config.Server) {
	return nil, global.G_Config
}

// @title    SetSystemConfig
// @description   set system config, 设置配置文件
// @auth                    （2020/04/05  20:22）
// @param     system         sys_model.System
// @return    err            error

func SetSystemConfig(system sys_model.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}
