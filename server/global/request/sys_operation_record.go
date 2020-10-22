package request

import "daigou-admin/model/sys_model"

type SysOperationRecordSearch struct {
	sys_model.SysOperationRecord
	PageInfo
}
