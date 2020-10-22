package response

import "daigou-admin/model/sys_model"

type ExaCustomerResponse struct {
	Customer sys_model.ExaCustomer `json:"customer"`
}
