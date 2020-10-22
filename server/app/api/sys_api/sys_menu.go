package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err, menus := sys_service.GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.SysMenusResponse{Menus: menus}, c)
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		respwrite.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, menuList, total := sys_service.GetInfoList()
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysBaseMenu true "新增菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func AddBaseMenu(c *gin.Context) {
	var menu sys_model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	MenuVerify := utils.Rules{
		"Path":      {utils.NotEmpty()},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0")},
	}
	MenuVerifyErr := utils.Verify(menu, MenuVerify)
	if MenuVerifyErr != nil {
		respwrite.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	MetaVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		respwrite.FailWithMessage(MetaVerifyErr.Error(), c)
		return
	}
	err := sys_service.AddBaseMenu(menu)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func GetBaseMenuTree(c *gin.Context) {
	err, menus := sys_service.GetBaseMenuTree()
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.SysBaseMenusResponse{Menus: menus}, c)
	}
}

// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func AddMenuAuthority(c *gin.Context) {
	var addMenuAuthorityInfo request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&addMenuAuthorityInfo)
	MenuVerify := utils.Rules{
		"AuthorityId": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(addMenuAuthorityInfo, MenuVerify)
	if MenuVerifyErr != nil {
		respwrite.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	err := sys_service.AddMenuAuthority(addMenuAuthorityInfo.Menus, addMenuAuthorityInfo.AuthorityId)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("添加成功", c)
	}
}

// @Tags authorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityIdInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func GetMenuAuthority(c *gin.Context) {
	var authorityIdInfo request.AuthorityIdInfo
	_ = c.ShouldBindJSON(&authorityIdInfo)
	MenuVerify := utils.Rules{
		"AuthorityId": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(authorityIdInfo, MenuVerify)
	if MenuVerifyErr != nil {
		respwrite.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	err, menus := sys_service.GetMenuAuthority(authorityIdInfo.AuthorityId)
	if err != nil {
		respwrite.FailWithDetailed(respwrite.ERROR, response.SysMenusResponse{Menus: menus}, fmt.Sprintf("添加失败，%v", err), c)
	} else {
		respwrite.Result(respwrite.SUCCESS, gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func DeleteBaseMenu(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		respwrite.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := sys_service.DeleteBaseMenu(idInfo.Id)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败：%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)

	}
}

// @Tags menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.SysBaseMenu true "更新菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updateBaseMenu [post]
func UpdateBaseMenu(c *gin.Context) {
	var menu sys_model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	MenuVerify := utils.Rules{
		"Path":      {"notEmpty"},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0")},
	}
	MenuVerifyErr := utils.Verify(menu, MenuVerify)
	if MenuVerifyErr != nil {
		respwrite.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	MetaVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		respwrite.FailWithMessage(MetaVerifyErr.Error(), c)
		return
	}
	err := sys_service.UpdateBaseMenu(menu)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("修改失败：%v", err), c)
	} else {
		respwrite.OkWithMessage("修改成功", c)
	}
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	MenuVerify := utils.Rules{
		"Id": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(idInfo, MenuVerify)
	if MenuVerifyErr != nil {
		respwrite.FailWithMessage(MenuVerifyErr.Error(), c)
		return
	}
	err, menu := sys_service.GetBaseMenuById(idInfo.Id)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询失败：%v", err), c)
	} else {
		respwrite.OkWithData(response.SysBaseMenuResponse{Menu: menu}, c)
	}
}
