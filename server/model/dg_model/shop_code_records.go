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


CREATE TABLE `shop_code_records` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '店铺ID(使用此店铺码开通的商店id)',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `bind_shop_id` bigint(20) unsigned NOT NULL COMMENT '生成店铺码的店铺ID',
  `gen_type` int(10) unsigned NOT NULL COMMENT '生成方式',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='店铺码'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// ShopCodeRecords struct is a row record of the shop_code_records table in the daigou database
type ShopCodeRecords struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID         uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID     uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`               // 店铺ID(使用此店铺码开通的商店id)//[ 2] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt  *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 创建时间//[ 3] bind_shop_id                                   ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	BindShopID uint64      `gorm:"column:bind_shop_id;type:ubigint;" json:"bind_shop_id"`     // 生成店铺码的店铺ID//[ 4] gen_type                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	GenType    uint32      `gorm:"column:gen_type;type:uint;" json:"gen_type"`                // 生成方式
}

var shop_code_recordsTableInfo = &TableInfo{
	Name: "shop_code_records",
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
			Name:               "shop_id",
			Comment:            `店铺ID(使用此店铺码开通的商店id)`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "uint64",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "bind_shop_id",
			Comment:            `生成店铺码的店铺ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "BindShopID",
			GoFieldType:        "uint64",
			JSONFieldName:      "bind_shop_id",
			ProtobufFieldName:  "bind_shop_id",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "gen_type",
			Comment:            `生成方式`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "GenType",
			GoFieldType:        "uint32",
			JSONFieldName:      "gen_type",
			ProtobufFieldName:  "gen_type",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopCodeRecords) TableName() string {
	return "shop_code_records"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopCodeRecords) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopCodeRecords) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopCodeRecords) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopCodeRecords) TableInfo() *TableInfo {
	return shop_code_recordsTableInfo
}

// GetAllShopCodeRecords is a function to get a slice of record(s) from shop_code_records table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopCodeRecords) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopCodeRecords, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopCodeRecords{})
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

// GetShopCodeRecords is a function to get a single record from the shop_code_records table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopCodeRecords) Get(DB *gorm.DB, argId uint32) (record *ShopCodeRecords, err error) {
	record = &ShopCodeRecords{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopCodeRecords is a function to add a single record to shop_code_records table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopCodeRecords) Add(DB *gorm.DB, record *ShopCodeRecords) (result *ShopCodeRecords, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopCodeRecords is a function to update a single record from shop_code_records table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopCodeRecords) Update(DB *gorm.DB, argId uint32, updated *ShopCodeRecords) (result *ShopCodeRecords, RowsAffected int64, err error) {

	result = &ShopCodeRecords{}
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

// DeleteShopCodeRecords is a function to delete a single record from shop_code_records table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopCodeRecords) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopCodeRecords{}
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
