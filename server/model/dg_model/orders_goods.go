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


CREATE TABLE `orders_goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单编号',
  `quantity` int(11) NOT NULL COMMENT '数量',
  `currency_type` int(11) NOT NULL DEFAULT '0' COMMENT '进价货币类型（0人民币1港币2澳门币3美元4英镑5欧元6韩元7日元）',
  `input_price` decimal(20,2) DEFAULT NULL COMMENT '进价',
  `single_price` decimal(20,2) NOT NULL COMMENT '销售单价',
  `total_input_price` decimal(20,2) DEFAULT NULL COMMENT '进价总价',
  `total_price` decimal(20,2) NOT NULL COMMENT '销售总价（单价*数量）',
  `specifications` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品规格（比如大小和颜色等）',
  `goods_name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品名称',
  `image` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '图片地址',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='订单对应商品表'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// OrdersGoods struct is a row record of the orders_goods table in the daigou database
type OrdersGoods struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID              uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`       // 主键//[ 1] order_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OrderID         string      `gorm:"column:order_id;type:varchar;" json:"order_id"`                   // 订单编号//[ 2] quantity                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Quantity        int32       `gorm:"column:quantity;type:int;" json:"quantity"`                       // 数量//[ 3] currency_type                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CurrencyType    int32       `gorm:"column:currency_type;type:int;default:0;" json:"currency_type"`   // 进价货币类型（0人民币1港币2澳门币3美元4英镑5欧元6韩元7日元）//[ 4] input_price                                    decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	InputPrice      float64     `gorm:"column:input_price;type:decimal;" json:"input_price"`             // 进价//[ 5] single_price                                   decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	SinglePrice     float64     `gorm:"column:single_price;type:decimal;" json:"single_price"`           // 销售单价//[ 6] total_input_price                              decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	TotalInputPrice float64     `gorm:"column:total_input_price;type:decimal;" json:"total_input_price"` // 进价总价//[ 7] total_price                                    decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	TotalPrice      float64     `gorm:"column:total_price;type:decimal;" json:"total_price"`             // 销售总价（单价*数量）//[ 8] specifications                                 varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Specifications  string      `gorm:"column:specifications;type:varchar;" json:"specifications"`       // 商品规格（比如大小和颜色等）//[ 9] goods_name                                     varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	GoodsName       string      `gorm:"column:goods_name;type:varchar;" json:"goods_name"`               // 商品名称//[10] image                                          varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Image           string      `gorm:"column:image;type:varchar;" json:"image"`                         // 图片地址//[11] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt       *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`              // 创建时间//[12] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt       *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`              // 更新时间//[13] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt       *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`              // 删除时间
}

var orders_goodsTableInfo = &TableInfo{
	Name: "orders_goods",
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
			Name:               "quantity",
			Comment:            `数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Quantity",
			GoFieldType:        "int32",
			JSONFieldName:      "quantity",
			ProtobufFieldName:  "quantity",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "currency_type",
			Comment:            `进价货币类型（0人民币1港币2澳门币3美元4英镑5欧元6韩元7日元）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CurrencyType",
			GoFieldType:        "int32",
			JSONFieldName:      "currency_type",
			ProtobufFieldName:  "currency_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "input_price",
			Comment:            `进价`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "InputPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "input_price",
			ProtobufFieldName:  "input_price",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "single_price",
			Comment:            `销售单价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "SinglePrice",
			GoFieldType:        "float64",
			JSONFieldName:      "single_price",
			ProtobufFieldName:  "single_price",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "total_input_price",
			Comment:            `进价总价`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "TotalInputPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "total_input_price",
			ProtobufFieldName:  "total_input_price",
			ProtobufType:       "float",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "total_price",
			Comment:            `销售总价（单价*数量）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "TotalPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "total_price",
			ProtobufFieldName:  "total_price",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "specifications",
			Comment:            `商品规格（比如大小和颜色等）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Specifications",
			GoFieldType:        "string",
			JSONFieldName:      "specifications",
			ProtobufFieldName:  "specifications",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "goods_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "GoodsName",
			GoFieldType:        "string",
			JSONFieldName:      "goods_name",
			ProtobufFieldName:  "goods_name",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "image",
			Comment:            `图片地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Image",
			GoFieldType:        "string",
			JSONFieldName:      "image",
			ProtobufFieldName:  "image",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrdersGoods) TableName() string {
	return "orders_goods"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrdersGoods) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrdersGoods) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrdersGoods) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrdersGoods) TableInfo() *TableInfo {
	return orders_goodsTableInfo
}

// GetAllOrdersGoods is a function to get a slice of record(s) from orders_goods table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (o *OrdersGoods) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*OrdersGoods, totalRows int64, err error) {

	resultOrm := DB.Model(&OrdersGoods{})
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

// GetOrdersGoods is a function to get a single record from the orders_goods table in the daigou database
// error - model.ErrNotFound, db Find error
func (o *OrdersGoods) Get(DB *gorm.DB, argId uint32) (record *OrdersGoods, err error) {
	record = &OrdersGoods{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOrdersGoods is a function to add a single record to orders_goods table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (o *OrdersGoods) Add(DB *gorm.DB, record *OrdersGoods) (result *OrdersGoods, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOrdersGoods is a function to update a single record from orders_goods table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (o *OrdersGoods) Update(DB *gorm.DB, argId uint32, updated *OrdersGoods) (result *OrdersGoods, RowsAffected int64, err error) {

	result = &OrdersGoods{}
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

// DeleteOrdersGoods is a function to delete a single record from orders_goods table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (o *OrdersGoods) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &OrdersGoods{}
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
