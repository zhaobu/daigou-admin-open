package initialize

import (
	_ "daigou-admin/docs"
	"daigou-admin/global"
	"daigou-admin/middleware"
	"daigou-admin/router"
	"daigou-admin/utils/zaplog"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	//Define format for the log of routes
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }
	gin.ForceConsoleColor()
	var Router = gin.Default()

	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	Router.Use(middleware.Ginzap(global.G_Logger.GetRequestLogger().Logger, zaplog.LogFormat, true))
	Router.Use(middleware.RecoveryWithZap(global.G_Logger.GetServerLogger().Logger, true))

	zaplog.Debug("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	zaplog.Debug("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	zaplog.Debug("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)                  // 注册用户路由
	router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
	router.InitApiRouter(ApiGroup)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
	router.InitSimpleUploaderRouter(ApiGroup)        // 断点续传（插件版）
	router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
	router.InitCasbinRouter(ApiGroup)                // 权限相关路由
	router.InitJwtRouter(ApiGroup)                   // jwt相关路由
	router.InitSystemRouter(ApiGroup)                // system相关路由
	router.InitCustomerRouter(ApiGroup)              // 客户路由
	router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
	router.InitSysDictionaryDetailRouter(ApiGroup)   // 字典详情管理
	router.InitSysDictionaryRouter(ApiGroup)         // 字典管理
	router.InitSysOperationRecordRouter(ApiGroup)    // 操作记录
	router.InitDaigouRouter(ApiGroup)                // 注册代购路由
	zaplog.Info("router register success")
	return Router
}
