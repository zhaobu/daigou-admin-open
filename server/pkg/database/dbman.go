package database

import (
	"daigou-admin/config"
	"daigou-admin/global"
	"daigou-admin/global/giface"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils/library"
	"daigou-admin/utils/zaplog"
)

type DBManage struct {
	giface.DBManageer
	sysDB *library.Mysql //管理后台数据库
	serDB *library.Mysql //业务数据库
}

func NewDBMange() *DBManage {
	return &DBManage{}
}

func (m *DBManage) Setup() {

	//admin数据库
	sysDB := global.G_Config.Database.Connects[0]
	sysDB.Type = config.SysDB
	m.Connect(sysDB, 0)

	//业务数据库
	serConf := global.G_Config.Service[0]
	serDB := &config.DBConnect{
		URL:  serConf.DB,
		Name: serConf.Name,
		Type: config.SerDB,
	}
	m.Connect(serDB, 0)

	global.G_DBManage = m
	zaplog.Info("initDBManager success")
}

func (m *DBManage) Connect(v *config.DBConnect, index int) (err error) {
	logConf := global.G_Config.Logger

	mysqlConf := &library.MysqlConf{
		DBConnect: v,
		Log:       true,
		LogPath:   logConf.Path + "/db",
		LogLevel:  logConf.Level,
		Index:     index,
	}

	if v.Type == config.SerDB { //连接业务数据库
		mysqlConf.SingularTable = true
		m.serDB, err = library.NewMysql(mysqlConf)
	} else if v.Type == config.SysDB { //连接管理后台数据库
		mysqlConf.SingularTable = false
		m.sysDB, err = library.NewMysql(mysqlConf)
	}
	if err != nil {
		zaplog.Errorf("Setup err:%s", err)
		return
	}
	return
}

func (m *DBManage) AutoMigrate() error {
	err := m.sysDB.AutoMigrate(
		sys_model.SysUser{},
		sys_model.SysAuthority{},
		sys_model.SysApi{},
		sys_model.SysBaseMenu{},
		sys_model.SysBaseMenuParameter{},
		sys_model.JwtBlacklist{},
		sys_model.SysWorkflow{},
		sys_model.SysWorkflowStepInfo{},
		sys_model.SysDictionary{},
		sys_model.SysDictionaryDetail{},
		sys_model.ExaFileUploadAndDownload{},
		sys_model.ExaFile{},
		sys_model.ExaFileChunk{},
		sys_model.ExaSimpleUploader{},
		sys_model.ExaCustomer{},
		sys_model.SysOperationRecord{}).Error
	if err != nil {
		zaplog.Errorf("AutoMigrate err:%s", err)
	}
	return err
}

func (m *DBManage) GetSysDB() *library.Mysql {
	return m.sysDB
}

func (m *DBManage) GetSerDB() *library.Mysql {
	return m.serDB
}

func (m *DBManage) Close() {
	m.serDB.Close()
	m.sysDB.Close()
}

//设置服务当前连接mysql
func (m *DBManage) ChangeSerDB(index int) (err error) {
	if m.serDB.Conf.Index == index {
		return
	}

	//关闭旧连接
	err = m.serDB.Close()
	if err != nil {
		return
	}

	serConfs := global.G_Config.Service

	//业务数据库
	serDB := &config.DBConnect{
		URL:  serConfs[index].DB,
		Name: serConfs[index].Name,
		Type: config.SerDB,
	}
	m.Connect(serDB, index)

	return
}
