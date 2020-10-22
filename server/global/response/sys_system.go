package response

import "daigou-admin/config"

type SysConfigResponse struct {
	Config *config.Server `json:"config"`
}
