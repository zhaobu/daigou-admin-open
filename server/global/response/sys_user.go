package response

import "daigou-admin/model/sys_model"

type SysUserResponse struct {
	User sys_model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      sys_model.SysUser `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
