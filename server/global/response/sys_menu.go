package response

import "daigou-admin/model/sys_model"

type SysMenusResponse struct {
	Menus []sys_model.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []sys_model.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu sys_model.SysBaseMenu `json:"menu"`
}
