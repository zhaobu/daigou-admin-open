package giface

import (
	"github.com/gin-gonic/gin"
)

// 路由
type Routerer interface {
	//初始化
	SetUp() *gin.Engine
}
