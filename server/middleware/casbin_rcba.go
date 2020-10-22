package middleware

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/utils/respwrite"
	"daigou-admin/utils/zaplog"

	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := sys_service.Casbin()
		res, err := e.Enforce(sub, obj, act)
		if err != nil {
			zaplog.Errorf("CasbinHandler err:%s", err)
			return
		}
		// 判断策略中是否存在
		// if global.G_Config.System.Mode == "debug" || res {
		if res {
			c.Next()
		} else {
			zaplog.Errorf("角色%s访问[%s]:%s权限不足", sub, act, obj)
			respwrite.Result(respwrite.ERROR, gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
