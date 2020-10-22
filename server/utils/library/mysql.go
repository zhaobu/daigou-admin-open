package library

import (
	"daigou-admin/config"
	"daigou-admin/utils/zaplog"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/os/gtimer"
	"github.com/jinzhu/gorm"
)

// Mysql连接器
type Mysql struct {
	*gorm.DB
	logger    *zaplog.Logger
	pingTimer *gtimer.Entry
	Conf      *MysqlConf
}

type DbGormLogger struct {
	*zaplog.Logger
}

type MysqlConf struct {
	*config.DBConnect
	SingularTable bool
	LogLevel      string
	LogPath       string
	Log           bool
	Index         int
}

func (logger *DbGormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logger.Sugar().Infof("file:%s time:[%s] sql:%s arg:%+v,index:%d", v[1], v[2], v[3], v[4], v[5])
	} else {
		logger.Sugar().Info(v...)
	}
	fmt.Println(gorm.LogFormatter(v...))
}

func newLogger(zlog *zaplog.Logger) *DbGormLogger {
	return &DbGormLogger{Logger: zlog}
}

// 连接数据库
func NewMysql(conf *MysqlConf) (res *Mysql, err error) {
	var (
		dbGorm *gorm.DB
		zlog   *zaplog.Logger
	)

	zaplog.Infof("DBLogger %s init success!", conf.Name)

	for i := 0; i < 60; i++ {
		zaplog.Infof("try to connect mysql %s the %d time", conf.URL, i)
		dbGorm, err = gorm.Open("mysql", conf.URL)
		err = dbGorm.DB().Ping() //测试是否连接成功
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		zaplog.Errorf("mysql connect error :", err)
		return
	}

	if conf.Log {
		//new数据库日志
		zlog = zaplog.NewLogger(&zaplog.Config{
			Logdir:    conf.LogPath,
			Debug:     false,
			LogName:   fmt.Sprintf("%s.log", conf.Name),
			Loglevel:  conf.LogLevel,
			AddCaller: false,
		}, 0)
		dbGorm.LogMode(true)
		dbGorm.SetLogger(newLogger(zlog))
	}

	zaplog.Infof("mysql %s connect success:%s", conf.Name, conf.URL)

	if dbGorm.Error != nil {
		zaplog.Panic(" database error :", dbGorm.Error)
	}
	//添加数据库心跳
	entry := gtimer.AddSingleton(time.Minute, func() {
		if err := dbGorm.DB().Ping(); err != nil {
			zaplog.Errorf("db %s pingGorm %s err=%s", conf.Name, conf.URL, err)
			return
		}
	})

	dbGorm.SingularTable(conf.SingularTable)
	sqlDb := dbGorm.DB()
	sqlDb.SetConnMaxLifetime(time.Minute * 10)

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)

	res = &Mysql{
		Conf:      conf,
		DB:        dbGorm,
		logger:    zlog,
		pingTimer: entry,
	}
	return
}

func (m *Mysql) Close() error {
	m.pingTimer.Close()
	return m.DB.Close()
}
