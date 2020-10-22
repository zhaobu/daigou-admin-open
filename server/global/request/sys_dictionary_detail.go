package request

import "daigou-admin/model/sys_model"

type SysDictionaryDetailSearch struct {
	sys_model.SysDictionaryDetail
	PageInfo
}
