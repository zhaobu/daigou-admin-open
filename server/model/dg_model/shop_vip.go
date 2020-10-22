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


CREATE TABLE `shop_vip` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) unsigned NOT NULL,
  `vip_level` int(10) unsigned NOT NULL COMMENT 'vip等级',
  `vip_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '等级名称',
  `opening_time` datetime NOT NULL COMMENT '开通时间',
  `end_time` datetime NOT NULL COMMENT '到期时间',
  `created_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='会员身份等级表'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// ShopVip struct is a row record of the shop_vip table in the daigou database
type ShopVip struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID          uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID      uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                  //[ 2] vip_level                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	VipLevel    uint32      `gorm:"column:vip_level;type:uint;" json:"vip_level"`              // vip等级//[ 3] vip_name                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	VipName     string      `gorm:"column:vip_name;type:varchar;" json:"vip_name"`             // 等级名称//[ 4] opening_time                                   datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	OpeningTime *model.Time `gorm:"column:opening_time;type:datetime;" json:"opening_time"`    // 开通时间//[ 5] end_time                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	EndTime     *model.Time `gorm:"column:end_time;type:datetime;" json:"end_time"`            // 到期时间//[ 6] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt   *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        //[ 7] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt   *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`        //[ 8] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt   *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var shop_vipTableInfo = &TableInfo{
	Name: "shop_vip",
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
			Name:               "user_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "vip_level",
			Comment:            `vip等级`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "VipLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "vip_level",
			ProtobufFieldName:  "vip_level",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "vip_name",
			Comment:            `等级名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "VipName",
			GoFieldType:        "string",
			JSONFieldName:      "vip_name",
			ProtobufFieldName:  "vip_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "opening_time",
			Comment:            `开通时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "OpeningTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "opening_time",
			ProtobufFieldName:  "opening_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "end_time",
			Comment:            `到期时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "created_at",
			Comment:            ``,
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
			Name:               "deleted_at",
			Comment:            ``,
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
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "updated_at",
			Comment:            ``,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopVip) TableName() string {
	return "shop_vip"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopVip) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopVip) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopVip) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopVip) TableInfo() *TableInfo {
	return shop_vipTableInfo
}

// GetAllShopVip is a function to get a slice of record(s) from shop_vip table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopVip) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopVip, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopVip{})
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

// GetShopVip is a function to get a single record from the shop_vip table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopVip) Get(DB *gorm.DB, argId uint32) (record *ShopVip, err error) {
	record = &ShopVip{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopVip is a function to add a single record to shop_vip table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopVip) Add(DB *gorm.DB, record *ShopVip) (result *ShopVip, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopVip is a function to update a single record from shop_vip table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopVip) Update(DB *gorm.DB, argId uint32, updated *ShopVip) (result *ShopVip, RowsAffected int64, err error) {

	result = &ShopVip{}
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

// DeleteShopVip is a function to delete a single record from shop_vip table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopVip) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopVip{}
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
