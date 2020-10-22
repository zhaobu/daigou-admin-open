package request

import "daigou-admin/model/sys_model"

type SysDictionarySearch struct {
	sys_model.SysDictionary
	PageInfo
}
