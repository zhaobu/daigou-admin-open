package request

import "daigou-admin/model/sys_model"

// api分页条件查询及排序结构体
type SearchApiParams struct {
	sys_model.SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
