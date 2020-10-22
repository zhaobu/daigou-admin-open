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


CREATE TABLE `system_ann` (
  `id` int(11) NOT NULL,
  `system_ann_version` varchar(10) COLLATE utf8mb4_bin NOT NULL COMMENT '版本号',
  `system_ann_target` int(10) unsigned NOT NULL COMMENT '0所有 1普通用户2代购',
  `system_ann_type` int(10) unsigned NOT NULL COMMENT '0拉取信息 1滚动展示 2弹框展示',
  `system_ann_msg` text COLLATE utf8mb4_bin NOT NULL COMMENT '内容',
  `system_ann_status` int(255) unsigned NOT NULL COMMENT '状态 0禁用 1开启',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='系统公告'


Comments
-------------------------------------
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 5] column is set for unsigned



*/

// SystemAnn struct is a row record of the system_ann table in the daigou database
type SystemAnn struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ID               int32       `gorm:"primary_key;column:id;type:int;" json:"id"`                         //[ 1] system_ann_version                             varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SystemAnnVersion string      `gorm:"column:system_ann_version;type:varchar;" json:"system_ann_version"` // 版本号//[ 2] system_ann_target                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	SystemAnnTarget  uint32      `gorm:"column:system_ann_target;type:uint;" json:"system_ann_target"`      // 0所有 1普通用户2代购//[ 3] system_ann_type                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	SystemAnnType    uint32      `gorm:"column:system_ann_type;type:uint;" json:"system_ann_type"`          // 0拉取信息 1滚动展示 2弹框展示//[ 4] system_ann_msg                                 text                 null: false  primary: false  isArray: false  auto: false  col: text            len: 0       default: []
	SystemAnnMsg     string      `gorm:"column:system_ann_msg;type:text;" json:"system_ann_msg"`            // 内容//[ 5] system_ann_status                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	SystemAnnStatus  uint32      `gorm:"column:system_ann_status;type:uint;" json:"system_ann_status"`      // 状态 0禁用 1开启//[ 6] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt        *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                // 创建时间//[ 7] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt        *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                // 更新时间//[ 8] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt        *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                // 删除时间
}

var system_annTableInfo = &TableInfo{
	Name: "system_ann",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},
		&ColumnInfo{
			Index:              1,
			Name:               "system_ann_version",
			Comment:            `版本号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SystemAnnVersion",
			GoFieldType:        "string",
			JSONFieldName:      "system_ann_version",
			ProtobufFieldName:  "system_ann_version",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "system_ann_target",
			Comment:            `0所有 1普通用户2代购`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SystemAnnTarget",
			GoFieldType:        "uint32",
			JSONFieldName:      "system_ann_target",
			ProtobufFieldName:  "system_ann_target",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "system_ann_type",
			Comment:            `0拉取信息 1滚动展示 2弹框展示`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SystemAnnType",
			GoFieldType:        "uint32",
			JSONFieldName:      "system_ann_type",
			ProtobufFieldName:  "system_ann_type",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "system_ann_msg",
			Comment:            `内容`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       0,
			GoFieldName:        "SystemAnnMsg",
			GoFieldType:        "string",
			JSONFieldName:      "system_ann_msg",
			ProtobufFieldName:  "system_ann_msg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "system_ann_status",
			Comment:            `状态 0禁用 1开启`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SystemAnnStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "system_ann_status",
			ProtobufFieldName:  "system_ann_status",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SystemAnn) TableName() string {
	return "system_ann"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SystemAnn) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SystemAnn) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SystemAnn) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SystemAnn) TableInfo() *TableInfo {
	return system_annTableInfo
}

// GetAllSystemAnn is a function to get a slice of record(s) from system_ann table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *SystemAnn) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*SystemAnn, totalRows int64, err error) {

	resultOrm := DB.Model(&SystemAnn{})
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

// GetSystemAnn is a function to get a single record from the system_ann table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *SystemAnn) Get(DB *gorm.DB, argId int32) (record *SystemAnn, err error) {
	record = &SystemAnn{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddSystemAnn is a function to add a single record to system_ann table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *SystemAnn) Add(DB *gorm.DB, record *SystemAnn) (result *SystemAnn, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateSystemAnn is a function to update a single record from system_ann table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *SystemAnn) Update(DB *gorm.DB, argId int32, updated *SystemAnn) (result *SystemAnn, RowsAffected int64, err error) {

	result = &SystemAnn{}
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

// DeleteSystemAnn is a function to delete a single record from system_ann table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *SystemAnn) Delete(DB *gorm.DB, argId int32) (rowsAffected int64, err error) {

	record := &SystemAnn{}
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
