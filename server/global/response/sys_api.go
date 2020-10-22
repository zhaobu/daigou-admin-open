package response

import "daigou-admin/model/sys_model"

type SysAPIResponse struct {
	Api sys_model.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []sys_model.SysApi `json:"apis"`
}
