package response

import "daigou-admin/model/sys_model"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File sys_model.ExaFile `json:"file"`
}
