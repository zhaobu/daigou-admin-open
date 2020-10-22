package giface

import (
	"daigou-admin/utils/zaplog"
)

// 日志
type Loggerer interface {
	//初始化
	Setup()
	GetServerLogger() *zaplog.Logger
	GetRequestLogger() *zaplog.Logger
}
