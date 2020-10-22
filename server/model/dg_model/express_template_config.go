package dg_model

import (
	"daigou-admin/model"

	"github.com/jinzhu/gorm"
)

var (
	_ *model.Time
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `express_template_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `templete_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '模板名称',
  `templete_data` text COLLATE utf8mb4_bin NOT NULL COMMENT '模板数据',
  `express_company_id` int(10) unsigned NOT NULL COMMENT '快递公司ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned



*/

// ExpressTemplateConfig struct is a row record of the express_template_config table in the daigou database
type ExpressTemplateConfig struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID               uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`      //[ 1] templete_name                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	TempleteName     string `gorm:"column:templete_name;type:varchar;" json:"templete_name"`        // 模板名称//[ 2] templete_data                                  text                 null: false  primary: false  isArray: false  auto: false  col: text            len: 0       default: []
	TempleteData     string `gorm:"column:templete_data;type:text;" json:"templete_data"`           // 模板数据//[ 3] express_company_id                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ExpressCompanyID uint32 `gorm:"column:express_company_id;type:uint;" json:"express_company_id"` // 快递公司ID
}

var express_template_configTableInfo = &TableInfo{
	Name: "express_template_config",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},
		&ColumnInfo{
			Index:              1,
			Name:               "templete_name",
			Comment:            `模板名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "TempleteName",
			GoFieldType:        "string",
			JSONFieldName:      "templete_name",
			ProtobufFieldName:  "templete_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "templete_data",
			Comment:            `模板数据`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       0,
			GoFieldName:        "TempleteData",
			GoFieldType:        "string",
			JSONFieldName:      "templete_data",
			ProtobufFieldName:  "templete_data",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "express_company_id",
			Comment:            `快递公司ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressCompanyID",
			GoFieldType:        "uint32",
			JSONFieldName:      "express_company_id",
			ProtobufFieldName:  "express_company_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *ExpressTemplateConfig) TableName() string {
	return "express_template_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *ExpressTemplateConfig) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *ExpressTemplateConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *ExpressTemplateConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *ExpressTemplateConfig) TableInfo() *TableInfo {
	return express_template_configTableInfo
}

// GetAllExpressTemplateConfig is a function to get a slice of record(s) from express_template_config table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (e *ExpressTemplateConfig) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ExpressTemplateConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&ExpressTemplateConfig{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = model.ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetExpressTemplateConfig is a function to get a single record from the express_template_config table in the daigou database
// error - model.ErrNotFound, db Find error
func (e *ExpressTemplateConfig) Get(DB *gorm.DB, argId uint32) (record *ExpressTemplateConfig, err error) {
	record = &ExpressTemplateConfig{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddExpressTemplateConfig is a function to add a single record to express_template_config table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (e *ExpressTemplateConfig) Add(DB *gorm.DB, record *ExpressTemplateConfig) (result *ExpressTemplateConfig, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateExpressTemplateConfig is a function to update a single record from express_template_config table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (e *ExpressTemplateConfig) Update(DB *gorm.DB, argId uint32, updated *ExpressTemplateConfig) (result *ExpressTemplateConfig, RowsAffected int64, err error) {

	result = &ExpressTemplateConfig{}
	db := DB.First(result, argId)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, model.ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteExpressTemplateConfig is a function to delete a single record from express_template_config table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (e *ExpressTemplateConfig) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ExpressTemplateConfig{}
	db := DB.First(record, argId)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
