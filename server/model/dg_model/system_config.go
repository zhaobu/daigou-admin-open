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


CREATE TABLE `system_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `class_id` int(10) unsigned NOT NULL COMMENT '字典编号',
  `class_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '字典说明',
  `class_value` json NOT NULL COMMENT '字典值',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='系统参数'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// SystemConfig struct is a row record of the system_config table in the daigou database
type SystemConfig struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID         uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] class_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ClassID    uint32 `gorm:"column:class_id;type:uint;" json:"class_id"`                // 字典编号//[ 2] class_name                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ClassName  string `gorm:"column:class_name;type:varchar;" json:"class_name"`         // 字典说明//[ 3] class_value                                    json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	ClassValue string `gorm:"column:class_value;type:json;" json:"class_value"`          // 字典值
}

var system_configTableInfo = &TableInfo{
	Name: "system_config",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
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
			Name:               "class_id",
			Comment:            `字典编号`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ClassID",
			GoFieldType:        "uint32",
			JSONFieldName:      "class_id",
			ProtobufFieldName:  "class_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "class_name",
			Comment:            `字典说明`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ClassName",
			GoFieldType:        "string",
			JSONFieldName:      "class_name",
			ProtobufFieldName:  "class_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "class_value",
			Comment:            `字典值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "ClassValue",
			GoFieldType:        "string",
			JSONFieldName:      "class_value",
			ProtobufFieldName:  "class_value",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SystemConfig) TableName() string {
	return "system_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SystemConfig) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SystemConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SystemConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SystemConfig) TableInfo() *TableInfo {
	return system_configTableInfo
}

// GetAllSystemConfig is a function to get a slice of record(s) from system_config table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *SystemConfig) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*SystemConfig, totalRows int64, err error) {

	resultOrm := DB.Model(&SystemConfig{})
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

// GetSystemConfig is a function to get a single record from the system_config table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *SystemConfig) Get(DB *gorm.DB, argId uint32) (record *SystemConfig, err error) {
	record = &SystemConfig{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddSystemConfig is a function to add a single record to system_config table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *SystemConfig) Add(DB *gorm.DB, record *SystemConfig) (result *SystemConfig, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateSystemConfig is a function to update a single record from system_config table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *SystemConfig) Update(DB *gorm.DB, argId uint32, updated *SystemConfig) (result *SystemConfig, RowsAffected int64, err error) {

	result = &SystemConfig{}
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

// DeleteSystemConfig is a function to delete a single record from system_config table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *SystemConfig) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &SystemConfig{}
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
