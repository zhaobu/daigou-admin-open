package response

import "daigou-admin/model/sys_model"

type ExaFileResponse struct {
	File sys_model.ExaFileUploadAndDownload `json:"file"`
}
