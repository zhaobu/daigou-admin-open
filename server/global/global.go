package global

import (
	"daigou-admin/config"
	"daigou-admin/global/giface"
	"daigou-admin/utils/library"

	"github.com/spf13/viper"
)

var (
	G_CacheManage giface.CacheManageer
	G_DBManage    giface.DBManageer
	G_Logger      giface.Loggerer
	G_Casbin      giface.Casbiner
	GVA_VP        *viper.Viper
	G_Config      *config.Server
)

func GetSysDB() *library.Mysql {
	return G_DBManage.GetSysDB()
}

func GetSerDB() *library.Mysql {
	return G_DBManage.GetSerDB()
}
