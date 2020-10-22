package router

import (
	"daigou-admin/app/api/dg_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDaigouRouter(Router *gin.RouterGroup) {
	group := Router.Group("daigou")
	group.Use(middleware.JWTAuth(), middleware.CasbinHandler())
	initServerConfig(group)
	initShopInfoRouter(group)
	initStaticsReportRouter(group)
	initUserInfoRouter(group)
	initProxyApi(group)
}

//代理访问
func initProxyApi(Router *gin.RouterGroup) {
	group := Router.Group("proxyApi")
	{
		group.POST("proxyPost", dg_api.ProxyPost) //代理post访问
	}
}

//服务器配置
func initServerConfig(Router *gin.RouterGroup) {
	group := Router.Group("serverConfig")
	{
		group.GET("getConnectList", dg_api.GetConnect)    //获取connect配置列表
		group.POST("changeConnect", dg_api.ChangeConnect) //修改connect连接
		// group.GET("getRedisList", dg_api.GetRedisList)    //获取redis配置列表
		// group.POST("changeRedis", dg_api.ChangeRedis)     //修改redis连接
		// group.GET("getMysqlList", dg_api.GetMysqlList)    //获取mysql配置列表
		// group.POST("changeMysql", dg_api.ChangeMysql)     //修改mysql连接
	}
}

//商店管理
func initShopInfoRouter(Router *gin.RouterGroup) {
	group := Router.Group("shopManage").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		group.DELETE("deleteShopInfoByIds", dg_api.DeleteShopInfoByIds) // 删除ShopInfo
		group.POST("findShopInfoList", dg_api.FindShopInfoList)         // 根据条件分页查找ShopInfo
		group.PUT("operateShop", dg_api.OperateShop)                    // 操作ShopInfo
	}
}

//用户管理
func initUserInfoRouter(Router *gin.RouterGroup) {
	group := Router.Group("userManage").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		group.DELETE("deleteUserInfoByIds", dg_api.DeleteUserInfoByIds) // 删除UserInfo
		group.POST("findUserInfoList", dg_api.FindUserInfoList)         // 根据条件分页查找UserInfo
		group.PUT("operateUser", dg_api.OperateUser)                    // 操作UserInfo
	}
}

//统计报表
func initStaticsReportRouter(Router *gin.RouterGroup) {
	group := Router.Group("staticsReport").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		group.GET("newUser", dg_api.NewUser)       // 查看新增用户
		group.GET("activeUser", dg_api.ActiveUser) // 查看活跃用户
	}
}
