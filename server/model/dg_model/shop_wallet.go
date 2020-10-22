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


CREATE TABLE `shop_wallet` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '商铺id',
  `month_orders` int(10) unsigned NOT NULL COMMENT '月订单个数',
  `month_profit` decimal(20,2) NOT NULL COMMENT '月收益',
  `month_cost` decimal(20,2) NOT NULL COMMENT '月成本',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='会员财富表'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// ShopWallet struct is a row record of the shop_wallet table in the daigou database
type ShopWallet struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID          uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID      uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`               // 商铺id//[ 2] month_orders                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	MonthOrders uint32      `gorm:"column:month_orders;type:uint;" json:"month_orders"`        // 月订单个数//[ 3] month_profit                                   decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	MonthProfit float64     `gorm:"column:month_profit;type:decimal;" json:"month_profit"`     // 月收益//[ 4] month_cost                                     decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	MonthCost   float64     `gorm:"column:month_cost;type:decimal;" json:"month_cost"`         // 月成本//[ 5] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt   *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        //[ 6] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt   *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var shop_walletTableInfo = &TableInfo{
	Name: "shop_wallet",
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
			Comment:            `商铺id`,
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
			Name:               "month_orders",
			Comment:            `月订单个数`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "MonthOrders",
			GoFieldType:        "uint32",
			JSONFieldName:      "month_orders",
			ProtobufFieldName:  "month_orders",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "month_profit",
			Comment:            `月收益`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "MonthProfit",
			GoFieldType:        "float64",
			JSONFieldName:      "month_profit",
			ProtobufFieldName:  "month_profit",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "month_cost",
			Comment:            `月成本`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "MonthCost",
			GoFieldType:        "float64",
			JSONFieldName:      "month_cost",
			ProtobufFieldName:  "month_cost",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopWallet) TableName() string {
	return "shop_wallet"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopWallet) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopWallet) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopWallet) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopWallet) TableInfo() *TableInfo {
	return shop_walletTableInfo
}

// GetAllShopWallet is a function to get a slice of record(s) from shop_wallet table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopWallet) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopWallet, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopWallet{})
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

// GetShopWallet is a function to get a single record from the shop_wallet table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopWallet) Get(DB *gorm.DB, argId uint32) (record *ShopWallet, err error) {
	record = &ShopWallet{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopWallet is a function to add a single record to shop_wallet table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopWallet) Add(DB *gorm.DB, record *ShopWallet) (result *ShopWallet, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopWallet is a function to update a single record from shop_wallet table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopWallet) Update(DB *gorm.DB, argId uint32, updated *ShopWallet) (result *ShopWallet, RowsAffected int64, err error) {

	result = &ShopWallet{}
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

// DeleteShopWallet is a function to delete a single record from shop_wallet table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopWallet) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopWallet{}
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
