package request

import "daigou-admin/model"

type {{.StructName}}Search struct{
    sys_model.{{.StructName}}
    PageInfo
}