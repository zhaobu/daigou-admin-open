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


CREATE TABLE `casbin_role_api` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '用户' COMMENT '角色名称',
  `api_list` json NOT NULL COMMENT '角色不能能访问的api列表',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE COMMENT '唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='角色api表'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// CasbinRoleAPI struct is a row record of the casbin_role_api table in the daigou database
type CasbinRoleAPI struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID        uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] name                                           varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: [用户]
	Name      string      `gorm:"column:name;type:varchar;default:用户;" json:"name"`          // 角色名称//[ 2] api_list                                       json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	APIList   string      `gorm:"column:api_list;type:json;" json:"api_list"`                // 角色不能能访问的api列表//[ 3] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 创建时间//[ 4] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`        // 更新时间//[ 5] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`        // 删除时间
}

var casbin_role_apiTableInfo = &TableInfo{
	Name: "casbin_role_api",
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
			Name:               "name",
			Comment:            `角色名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "api_list",
			Comment:            `角色不能能访问的api列表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "APIList",
			GoFieldType:        "string",
			JSONFieldName:      "api_list",
			ProtobufFieldName:  "api_list",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "created_at",
			Comment:            `创建时间`,
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
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "updated_at",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "deleted_at",
			Comment:            `删除时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "DeletedAt",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "deleted_at",
			ProtobufFieldName:  "deleted_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CasbinRoleAPI) TableName() string {
	return "casbin_role_api"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CasbinRoleAPI) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CasbinRoleAPI) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CasbinRoleAPI) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CasbinRoleAPI) TableInfo() *TableInfo {
	return casbin_role_apiTableInfo
}

// GetAllCasbinRoleAPI is a function to get a slice of record(s) from casbin_role_api table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (c *CasbinRoleAPI) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*CasbinRoleAPI, totalRows int64, err error) {

	resultOrm := DB.Model(&CasbinRoleAPI{})
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

// GetCasbinRoleAPI is a function to get a single record from the casbin_role_api table in the daigou database
// error - model.ErrNotFound, db Find error
func (c *CasbinRoleAPI) Get(DB *gorm.DB, argId uint32) (record *CasbinRoleAPI, err error) {
	record = &CasbinRoleAPI{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddCasbinRoleAPI is a function to add a single record to casbin_role_api table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (c *CasbinRoleAPI) Add(DB *gorm.DB, record *CasbinRoleAPI) (result *CasbinRoleAPI, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCasbinRoleAPI is a function to update a single record from casbin_role_api table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (c *CasbinRoleAPI) Update(DB *gorm.DB, argId uint32, updated *CasbinRoleAPI) (result *CasbinRoleAPI, RowsAffected int64, err error) {

	result = &CasbinRoleAPI{}
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

// DeleteCasbinRoleAPI is a function to delete a single record from casbin_role_api table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (c *CasbinRoleAPI) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &CasbinRoleAPI{}
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
