package request

import "daigou-admin/model/sys_model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []sys_model.SysBaseMenu
	AuthorityId string
}

// Get role by id structure
type AuthorityIdInfo struct {
	AuthorityId string
}
