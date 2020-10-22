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


CREATE TABLE `third_party_interface` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `third_party_id` int(11) NOT NULL COMMENT '第三方接口平台id',
  `third_name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '第三方接口名称',
  `third_url` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '第三方接口url链接',
  `created_at` datetime DEFAULT NULL COMMENT '编辑时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='第三方快递接口'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// ThirdPartyInterface struct is a row record of the third_party_interface table in the daigou database
type ThirdPartyInterface struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID           uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] third_party_id                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ThirdPartyID int32       `gorm:"column:third_party_id;type:int;" json:"third_party_id"`     // 第三方接口平台id//[ 2] third_name                                     varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ThirdName    string      `gorm:"column:third_name;type:varchar;" json:"third_name"`         // 第三方接口名称//[ 3] third_url                                      varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ThirdURL     string      `gorm:"column:third_url;type:varchar;" json:"third_url"`           // 第三方接口url链接//[ 4] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt    *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 编辑时间
}

var third_party_interfaceTableInfo = &TableInfo{
	Name: "third_party_interface",
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
			Name:               "third_party_id",
			Comment:            `第三方接口平台id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ThirdPartyID",
			GoFieldType:        "int32",
			JSONFieldName:      "third_party_id",
			ProtobufFieldName:  "third_party_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "third_name",
			Comment:            `第三方接口名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ThirdName",
			GoFieldType:        "string",
			JSONFieldName:      "third_name",
			ProtobufFieldName:  "third_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "third_url",
			Comment:            `第三方接口url链接`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ThirdURL",
			GoFieldType:        "string",
			JSONFieldName:      "third_url",
			ProtobufFieldName:  "third_url",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "created_at",
			Comment:            `编辑时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *ThirdPartyInterface) TableName() string {
	return "third_party_interface"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *ThirdPartyInterface) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *ThirdPartyInterface) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *ThirdPartyInterface) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *ThirdPartyInterface) TableInfo() *TableInfo {
	return third_party_interfaceTableInfo
}

// GetAllThirdPartyInterface is a function to get a slice of record(s) from third_party_interface table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (t *ThirdPartyInterface) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ThirdPartyInterface, totalRows int64, err error) {

	resultOrm := DB.Model(&ThirdPartyInterface{})
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

// GetThirdPartyInterface is a function to get a single record from the third_party_interface table in the daigou database
// error - model.ErrNotFound, db Find error
func (t *ThirdPartyInterface) Get(DB *gorm.DB, argId uint32) (record *ThirdPartyInterface, err error) {
	record = &ThirdPartyInterface{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddThirdPartyInterface is a function to add a single record to third_party_interface table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (t *ThirdPartyInterface) Add(DB *gorm.DB, record *ThirdPartyInterface) (result *ThirdPartyInterface, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateThirdPartyInterface is a function to update a single record from third_party_interface table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (t *ThirdPartyInterface) Update(DB *gorm.DB, argId uint32, updated *ThirdPartyInterface) (result *ThirdPartyInterface, RowsAffected int64, err error) {

	result = &ThirdPartyInterface{}
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

// DeleteThirdPartyInterface is a function to delete a single record from third_party_interface table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (t *ThirdPartyInterface) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ThirdPartyInterface{}
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
