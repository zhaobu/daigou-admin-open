package response

import "daigou-admin/model/sys_model"

type SysAuthorityResponse struct {
	Authority sys_model.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      sys_model.SysAuthority `json:"authority"`
	OldAuthorityId string                 `json:"oldAuthorityId"`
}
