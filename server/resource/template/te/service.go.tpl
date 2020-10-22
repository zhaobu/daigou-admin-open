package service

import (
	"daigou-admin/global"
	"daigou-admin/model"
	"daigou-admin/global/request"
)

// @title    Create{{.StructName}}
// @description   create a {{.StructName}}
// @param     {{.Abbreviation}}               sys_model.{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @return    err             error

func Create{{.StructName}}({{.Abbreviation}} sys_model.{{.StructName}}) (err error) {
	err =global.GetSysDB(),Create(&{{.Abbreviation}}).Error
	return err
}

// @title    Delete{{.StructName}}
// @description   delete a {{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     {{.Abbreviation}}               sys_model.{{.StructName}}
// @return                    error

func Delete{{.StructName}}({{.Abbreviation}} sys_model.{{.StructName}}) (err error) {
	err =global.GetSysDB(),Delete({{.Abbreviation}}).Error
	return err
}

// @title    Delete{{.StructName}}ByIds
// @description   delete {{.StructName}}s
// @auth                     （2020/04/05  20:22）
// @param     {{.Abbreviation}}               sys_model.{{.StructName}}
// @return                    error

func Delete{{.StructName}}ByIds(ids request.IdsReq) (err error) {
	err =global.GetSysDB(),Delete(&[]sys_model.{{.StructName}}{},"id in (?)",ids.Ids).Error
	return err
}

// @title    Update{{.StructName}}
// @description   update a {{.StructName}}
// @param     {{.Abbreviation}}          *sys_model.{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @return                    error

func Update{{.StructName}}({{.Abbreviation}} *sys_model.{{.StructName}}) (err error) {
	err =global.GetSysDB(),Save({{.Abbreviation}}).Error
	return err
}

// @title    Get{{.StructName}}
// @description   get the info of a {{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    {{.StructName}}        {{.StructName}}

func Get{{.StructName}}(id uint) (err error, {{.Abbreviation}} sys_model.{{.StructName}}) {
	err =global.GetSysDB(),Where("id = ?", id).First(&{{.Abbreviation}}).Error
	return
}

// @title    Get{{.StructName}}InfoList
// @description   get {{.StructName}} list by pagination, 分页获取{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func Get{{.StructName}}InfoList(info request.{{.StructName}}Search) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db :=global.GetSysDB(),Model(&sys_model.{{.StructName}}{})
    var {{.Abbreviation}}s []sys_model.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if info.{{.FieldName}} != 0 {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != 0 {
        db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if !info.{{.FieldName}}.IsZero() {
         db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- end }}
        {{- end }}
    {{- end }}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
	return err, {{.Abbreviation}}s, total
}