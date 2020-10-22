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


CREATE TABLE `shop_vip_price` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '商品名称',
  `member_discount_price` decimal(10,2) unsigned NOT NULL COMMENT '商品折扣价格',
  `member_original_price` decimal(10,2) NOT NULL COMMENT '商品原始价格',
  `member_jian_shao` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '商品简绍',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned



*/

// ShopVipPrice struct is a row record of the shop_vip_price table in the daigou database
type ShopVipPrice struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID                  uint32  `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`                //[ 1] member_name                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	MemberName          string  `gorm:"column:member_name;type:varchar;" json:"member_name"`                      // 商品名称//[ 2] member_discount_price                          udecimal             null: false  primary: false  isArray: false  auto: false  col: udecimal        len: -1      default: []
	MemberDiscountPrice float64 `gorm:"column:member_discount_price;type:udecimal;" json:"member_discount_price"` // 商品折扣价格//[ 3] member_original_price                          decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	MemberOriginalPrice float64 `gorm:"column:member_original_price;type:decimal;" json:"member_original_price"`  // 商品原始价格//[ 4] member_jian_shao                               varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	MemberJianShao      string  `gorm:"column:member_jian_shao;type:varchar;" json:"member_jian_shao"`            // 商品简绍
}

var shop_vip_priceTableInfo = &TableInfo{
	Name: "shop_vip_price",
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
			Name:               "member_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "MemberName",
			GoFieldType:        "string",
			JSONFieldName:      "member_name",
			ProtobufFieldName:  "member_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "member_discount_price",
			Comment:            `商品折扣价格`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "udecimal",
			DatabaseTypePretty: "udecimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "udecimal",
			ColumnLength:       -1,
			GoFieldName:        "MemberDiscountPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "member_discount_price",
			ProtobufFieldName:  "member_discount_price",
			ProtobufType:       "float",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "member_original_price",
			Comment:            `商品原始价格`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "MemberOriginalPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "member_original_price",
			ProtobufFieldName:  "member_original_price",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "member_jian_shao",
			Comment:            `商品简绍`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "MemberJianShao",
			GoFieldType:        "string",
			JSONFieldName:      "member_jian_shao",
			ProtobufFieldName:  "member_jian_shao",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopVipPrice) TableName() string {
	return "shop_vip_price"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopVipPrice) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopVipPrice) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopVipPrice) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopVipPrice) TableInfo() *TableInfo {
	return shop_vip_priceTableInfo
}

// GetAllShopVipPrice is a function to get a slice of record(s) from shop_vip_price table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopVipPrice) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopVipPrice, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopVipPrice{})
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

// GetShopVipPrice is a function to get a single record from the shop_vip_price table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopVipPrice) Get(DB *gorm.DB, argId uint32) (record *ShopVipPrice, err error) {
	record = &ShopVipPrice{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopVipPrice is a function to add a single record to shop_vip_price table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopVipPrice) Add(DB *gorm.DB, record *ShopVipPrice) (result *ShopVipPrice, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopVipPrice is a function to update a single record from shop_vip_price table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopVipPrice) Update(DB *gorm.DB, argId uint32, updated *ShopVipPrice) (result *ShopVipPrice, RowsAffected int64, err error) {

	result = &ShopVipPrice{}
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

// DeleteShopVipPrice is a function to delete a single record from shop_vip_price table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopVipPrice) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopVipPrice{}
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
