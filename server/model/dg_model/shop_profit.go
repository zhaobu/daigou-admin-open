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


CREATE TABLE `shop_profit` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单编号',
  `user_id` int(10) unsigned NOT NULL COMMENT '买家用户编号',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '卖家用户编号(或称店铺编号)',
  `price` decimal(20,2) NOT NULL COMMENT '总金额',
  `profit` decimal(20,2) NOT NULL COMMENT '订单利润',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// ShopProfit struct is a row record of the shop_profit table in the daigou database
type ShopProfit struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID        uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] order_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OrderID   string      `gorm:"column:order_id;type:varchar;" json:"order_id"`             // 订单编号//[ 2] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID    uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                  // 买家用户编号//[ 3] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID    uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`               // 卖家用户编号(或称店铺编号)//[ 4] price                                          decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Price     float64     `gorm:"column:price;type:decimal;" json:"price"`                   // 总金额//[ 5] profit                                         decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Profit    float64     `gorm:"column:profit;type:decimal;" json:"profit"`                 // 订单利润//[ 6] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 创建时间
}

var shop_profitTableInfo = &TableInfo{
	Name: "shop_profit",
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
			Name:               "order_id",
			Comment:            `订单编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "OrderID",
			GoFieldType:        "string",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "user_id",
			Comment:            `买家用户编号`,
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
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "shop_id",
			Comment:            `卖家用户编号(或称店铺编号)`,
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
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "price",
			Comment:            `总金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "Price",
			GoFieldType:        "float64",
			JSONFieldName:      "price",
			ProtobufFieldName:  "price",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "profit",
			Comment:            `订单利润`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "Profit",
			GoFieldType:        "float64",
			JSONFieldName:      "profit",
			ProtobufFieldName:  "profit",
			ProtobufType:       "float",
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
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopProfit) TableName() string {
	return "shop_profit"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopProfit) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopProfit) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopProfit) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopProfit) TableInfo() *TableInfo {
	return shop_profitTableInfo
}

// GetAllShopProfit is a function to get a slice of record(s) from shop_profit table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopProfit) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopProfit, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopProfit{})
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

// GetShopProfit is a function to get a single record from the shop_profit table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopProfit) Get(DB *gorm.DB, argId uint32) (record *ShopProfit, err error) {
	record = &ShopProfit{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopProfit is a function to add a single record to shop_profit table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopProfit) Add(DB *gorm.DB, record *ShopProfit) (result *ShopProfit, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopProfit is a function to update a single record from shop_profit table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopProfit) Update(DB *gorm.DB, argId uint32, updated *ShopProfit) (result *ShopProfit, RowsAffected int64, err error) {

	result = &ShopProfit{}
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

// DeleteShopProfit is a function to delete a single record from shop_profit table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopProfit) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopProfit{}
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
