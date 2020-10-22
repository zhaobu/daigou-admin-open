package sys_service

import (
	"daigou-admin/global"
	"daigou-admin/model/sys_model"
	"errors"
	"strconv"
)

// @title   getMenuTreeMap
// @description    获取路由总树map
// @auth       qm      (2020/05/06 10:26)
// @return     err             error
// @return    menusMsp            map{string}[]SysBaseMenu

func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]sys_model.SysMenu) {
	var allMenus []sys_model.SysMenu
	treeMap = make(map[string][]sys_model.SysMenu)
	err = global.GetSysDB().Where("authority_id = ?", authorityId).Order("sort", true).Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// @title    GetMenuTree
// @description   获取动态菜单树
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return    err             error
// @return    menus           []sys_model.SysMenu

func GetMenuTree(authorityId string) (err error, menus []sys_model.SysMenu) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// @title    getChildrenList
// @description   获取子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *sys_model.SysMenu
// @param     sql             string
// @return    err             error

func getChildrenList(menu *sys_model.SysMenu, treeMap map[string][]sys_model.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// @title    GetInfoList
// @description   获取路由分页
// @auth                     （2020/04/05  20:22）
// @param     info            request.PageInfo
// @return    err             error
// @return    list            interface{}
// @return    total           int

func GetInfoList() (err error, list interface{}, total int) {
	var menuList []sys_model.SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

// @title    getBaseChildrenList
// @description   get children of menu, 获取菜单的子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *sys_model.SysBaseMenu
// @return    err             error

func getBaseChildrenList(menu *sys_model.SysBaseMenu, treeMap map[string][]sys_model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// @title    AddBaseMenu
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22）
// @param     menu            *sys_model.SysBaseMenu
// @return    err             error
// 增加基础路由

func AddBaseMenu(menu sys_model.SysBaseMenu) (err error) {
	findOne := global.GetSysDB().Where("name = ?", menu.Name).Find(&sys_model.SysBaseMenu{}).Error
	if findOne != nil {
		err = global.GetSysDB().Create(&menu).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

// @title   getBaseMenuTreeMap
// @description    获取路由总树map
// @auth       qm      (2020/05/06 10:26)
// @return     err             error
// @return    menusMsp            map{string}[]SysBaseMenu

func getBaseMenuTreeMap() (err error, treeMap map[string][]sys_model.SysBaseMenu) {
	var allMenus []sys_model.SysBaseMenu
	treeMap = make(map[string][]sys_model.SysBaseMenu)
	err = global.GetSysDB().Order("sort", true).Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// @title    GetBaseMenuTree
// @description   获取基础路由树
// @auth                     （2020/04/05  20:22）
// @return    err              error
// @return    menus            []SysBaseMenu

func GetBaseMenuTree() (err error, menus []sys_model.SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

// @title    AddMenuAuthority
// @description   为角色增加menu树
// @auth                     （2020/04/05  20:22）
// @param     menus           []sys_model.SysBaseMenu
// @param     authorityId     string
// @return                    error

func AddMenuAuthority(menus []sys_model.SysBaseMenu, authorityId string) (err error) {
	var auth sys_model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

// @title    GetMenuAuthority
// @description   查看当前角色树
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return    err             error
// @return    menus           []SysBaseMenu

func GetMenuAuthority(authorityId string) (err error, menus []sys_model.SysMenu) {
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	err = global.GetSysDB().Order("sort", true).Find(&menus).Error
	//err = global.GetSysDB().Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
