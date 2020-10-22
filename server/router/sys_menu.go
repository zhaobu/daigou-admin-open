package router

import (
	"daigou-admin/app/api/sys_api"
	"daigou-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		MenuRouter.POST("getMenu", sys_api.GetMenu)                   // 获取菜单树
		MenuRouter.POST("getMenuList", sys_api.GetMenuList)           // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", sys_api.AddBaseMenu)           // 新增菜单
		MenuRouter.POST("getBaseMenuTree", sys_api.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", sys_api.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", sys_api.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", sys_api.DeleteBaseMenu)     // 删除菜单
		MenuRouter.POST("updateBaseMenu", sys_api.UpdateBaseMenu)     // 更新菜单
		MenuRouter.POST("getBaseMenuById", sys_api.GetBaseMenuById)   // 根据id获取菜单
	}
	return MenuRouter
}
