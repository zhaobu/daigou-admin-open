package response

import "daigou-admin/global/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
