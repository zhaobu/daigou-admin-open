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


CREATE TABLE `system_express_company` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `express_code` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '快递编码(快递鸟)',
  `express_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '快递公司名称',
  `exoress_code_kd100` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '快递编码(快递100)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='系统快递公司资料'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// SystemExpressCompany struct is a row record of the system_express_company table in the daigou database
type SystemExpressCompany struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID               uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`          // 主键//[ 1] express_code                                   varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ExpressCode      string `gorm:"column:express_code;type:varchar;" json:"express_code"`              // 快递编码(快递鸟)//[ 2] express_name                                   varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ExpressName      string `gorm:"column:express_name;type:varchar;" json:"express_name"`              // 快递公司名称//[ 3] exoress_code_kd100                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ExoressCodeKd100 string `gorm:"column:exoress_code_kd100;type:varchar;" json:"exoress_code_kd_100"` // 快递编码(快递100)
}

var system_express_companyTableInfo = &TableInfo{
	Name: "system_express_company",
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
			Name:               "express_code",
			Comment:            `快递编码(快递鸟)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ExpressCode",
			GoFieldType:        "string",
			JSONFieldName:      "express_code",
			ProtobufFieldName:  "express_code",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "express_name",
			Comment:            `快递公司名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ExpressName",
			GoFieldType:        "string",
			JSONFieldName:      "express_name",
			ProtobufFieldName:  "express_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "exoress_code_kd100",
			Comment:            `快递编码(快递100)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ExoressCodeKd100",
			GoFieldType:        "string",
			JSONFieldName:      "exoress_code_kd_100",
			ProtobufFieldName:  "exoress_code_kd_100",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SystemExpressCompany) TableName() string {
	return "system_express_company"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SystemExpressCompany) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SystemExpressCompany) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SystemExpressCompany) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SystemExpressCompany) TableInfo() *TableInfo {
	return system_express_companyTableInfo
}

// GetAllSystemExpressCompany is a function to get a slice of record(s) from system_express_company table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *SystemExpressCompany) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*SystemExpressCompany, totalRows int64, err error) {

	resultOrm := DB.Model(&SystemExpressCompany{})
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

// GetSystemExpressCompany is a function to get a single record from the system_express_company table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *SystemExpressCompany) Get(DB *gorm.DB, argId uint32) (record *SystemExpressCompany, err error) {
	record = &SystemExpressCompany{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddSystemExpressCompany is a function to add a single record to system_express_company table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *SystemExpressCompany) Add(DB *gorm.DB, record *SystemExpressCompany) (result *SystemExpressCompany, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateSystemExpressCompany is a function to update a single record from system_express_company table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *SystemExpressCompany) Update(DB *gorm.DB, argId uint32, updated *SystemExpressCompany) (result *SystemExpressCompany, RowsAffected int64, err error) {

	result = &SystemExpressCompany{}
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

// DeleteSystemExpressCompany is a function to delete a single record from system_express_company table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *SystemExpressCompany) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &SystemExpressCompany{}
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
