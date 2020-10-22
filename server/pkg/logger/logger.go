package logger

import (
	"daigou-admin/global"
	"daigou-admin/global/giface"
	"daigou-admin/utils/zaplog"
)

type Logger struct {
	giface.Loggerer
	RequestLogger *zaplog.Logger
	ServerLogger  *zaplog.Logger
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) GetServerLogger() *zaplog.Logger {
	return l.ServerLogger
}

func (l *Logger) GetRequestLogger() *zaplog.Logger {
	return l.RequestLogger
}

func (l *Logger) Setup() {
	logConf := global.G_Config.Logger
	if logConf.EnabledBUS {
		l.ServerLogger = zaplog.InitLog(&zaplog.Config{
			Logdir:    logConf.Path + "/bus",
			Debug:     logConf.Stdout,
			LogName:   "server.log",
			Loglevel:  logConf.Level,
			AddCaller: true,
		})
		zaplog.Info("Logger init success!")
	}

	l.RequestLogger = zaplog.NewLogger(&zaplog.Config{
		Logdir:    logConf.Path + "/request",
		Debug:     false,
		LogName:   "access.log",
		Loglevel:  logConf.Level,
		AddCaller: false,
	}, 0)
	zaplog.Info("RequestLogger init success!")

	global.G_Logger = l
}
