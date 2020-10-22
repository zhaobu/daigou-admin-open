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


CREATE TABLE `orders_printing` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `shop_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '店铺编号',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单编号',
  `printing_count` int(10) unsigned NOT NULL COMMENT '打印次数',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned



*/

// OrdersPrinting struct is a row record of the orders_printing table in the daigou database
type OrdersPrinting struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID            uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ShopID        uint64      `gorm:"column:shop_id;type:ubigint;default:0;" json:"shop_id"`     // 店铺编号//[ 2] order_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OrderID       string      `gorm:"column:order_id;type:varchar;" json:"order_id"`             // 订单编号//[ 3] printing_count                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	PrintingCount uint32      `gorm:"column:printing_count;type:uint;" json:"printing_count"`    // 打印次数//[ 4] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt     *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 创建时间//[ 5] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt     *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`        // 更新时间
}

var orders_printingTableInfo = &TableInfo{
	Name: "orders_printing",
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
			Comment:            `店铺编号`,
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
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "printing_count",
			Comment:            `打印次数`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PrintingCount",
			GoFieldType:        "uint32",
			JSONFieldName:      "printing_count",
			ProtobufFieldName:  "printing_count",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrdersPrinting) TableName() string {
	return "orders_printing"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrdersPrinting) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrdersPrinting) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrdersPrinting) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrdersPrinting) TableInfo() *TableInfo {
	return orders_printingTableInfo
}

// GetAllOrdersPrinting is a function to get a slice of record(s) from orders_printing table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (o *OrdersPrinting) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*OrdersPrinting, totalRows int64, err error) {

	resultOrm := DB.Model(&OrdersPrinting{})
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

// GetOrdersPrinting is a function to get a single record from the orders_printing table in the daigou database
// error - model.ErrNotFound, db Find error
func (o *OrdersPrinting) Get(DB *gorm.DB, argId uint32) (record *OrdersPrinting, err error) {
	record = &OrdersPrinting{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOrdersPrinting is a function to add a single record to orders_printing table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (o *OrdersPrinting) Add(DB *gorm.DB, record *OrdersPrinting) (result *OrdersPrinting, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOrdersPrinting is a function to update a single record from orders_printing table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (o *OrdersPrinting) Update(DB *gorm.DB, argId uint32, updated *OrdersPrinting) (result *OrdersPrinting, RowsAffected int64, err error) {

	result = &OrdersPrinting{}
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

// DeleteOrdersPrinting is a function to delete a single record from orders_printing table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (o *OrdersPrinting) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &OrdersPrinting{}
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
