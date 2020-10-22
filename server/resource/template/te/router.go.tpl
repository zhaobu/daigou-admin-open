package router

import (
	"daigou-admin/api/v1"
	"daigou-admin/middleware"
	"github.com/gin-gonic/gin"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.StructName}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		{{.StructName}}Router.POST("create{{.StructName}}", sys_api.Create{{.StructName}})   // 新建{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}", sys_api.Delete{{.StructName}}) // 删除{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}ByIds", sys_api.Delete{{.StructName}}ByIds) // 批量删除{{.StructName}}
		{{.StructName}}Router.PUT("update{{.StructName}}", sys_api.Update{{.StructName}})    // 更新{{.StructName}}
		{{.StructName}}Router.GET("find{{.StructName}}", sys_api.Find{{.StructName}})        // 根据ID获取{{.StructName}}
		{{.StructName}}Router.GET("get{{.StructName}}List", sys_api.Get{{.StructName}}List)  // 获取{{.StructName}}列表
	}
}
