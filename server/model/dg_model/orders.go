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


CREATE TABLE `orders` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单编号',
  `user_id` int(10) unsigned NOT NULL COMMENT '买家用户编号',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '卖家用户编号(或称店铺编号)',
  `order_status` int(11) NOT NULL COMMENT '订单状态0待确认1待发货2已发货3已完成',
  `preferential_price` decimal(20,2) DEFAULT NULL COMMENT '优惠金额',
  `price` decimal(20,2) NOT NULL COMMENT '总金额',
  `profit` decimal(20,2) NOT NULL COMMENT '订单利润',
  `goods_name_arr` json DEFAULT NULL COMMENT '订单商品列表',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '买家备注',
  `ispay` int(11) NOT NULL COMMENT '是否已支付',
  `pay_price` decimal(20,2) NOT NULL COMMENT '收取金额',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '此订单状态1正常0禁用-1删除',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `ship_time` datetime DEFAULT NULL COMMENT '发货时间',
  `complete_time` datetime DEFAULT NULL COMMENT '完成时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `confirm_time` datetime DEFAULT NULL COMMENT '确认时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_id` (`order_id`) USING BTREE,
  KEY `shop_id` (`shop_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='订单表'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// Orders struct is a row record of the orders table in the daigou database
type Orders struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID                uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`         // 主键//[ 1] order_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OrderID           string      `gorm:"column:order_id;type:varchar;" json:"order_id"`                     // 订单编号//[ 2] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID            uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                          // 买家用户编号//[ 3] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID            uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`                       // 卖家用户编号(或称店铺编号)//[ 4] order_status                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	OrderStatus       int32       `gorm:"column:order_status;type:int;" json:"order_status"`                 // 订单状态0待确认1待发货2已发货3已完成//[ 5] preferential_price                             decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	PreferentialPrice float64     `gorm:"column:preferential_price;type:decimal;" json:"preferential_price"` // 优惠金额//[ 6] price                                          decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Price             float64     `gorm:"column:price;type:decimal;" json:"price"`                           // 总金额//[ 7] profit                                         decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Profit            float64     `gorm:"column:profit;type:decimal;" json:"profit"`                         // 订单利润//[ 8] goods_name_arr                                 json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	GoodsNameArr      string      `gorm:"column:goods_name_arr;type:json;" json:"goods_name_arr"`            // 订单商品列表//[ 9] remark                                         varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Remark            string      `gorm:"column:remark;type:varchar;" json:"remark"`                         // 买家备注//[10] ispay                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Ispay             int32       `gorm:"column:ispay;type:int;" json:"ispay"`                               // 是否已支付//[11] pay_price                                      decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	PayPrice          float64     `gorm:"column:pay_price;type:decimal;" json:"pay_price"`                   // 收取金额//[12] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Status            int32       `gorm:"column:status;type:int;default:1;" json:"status"`                   // 此订单状态1正常0禁用-1删除//[13] pay_time                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	PayTime           *model.Time `gorm:"column:pay_time;type:datetime;" json:"pay_time"`                    // 支付时间//[14] ship_time                                      datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ShipTime          *model.Time `gorm:"column:ship_time;type:datetime;" json:"ship_time"`                  // 发货时间//[15] complete_time                                  datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CompleteTime      *model.Time `gorm:"column:complete_time;type:datetime;" json:"complete_time"`          // 完成时间//[16] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt         *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                // 创建时间//[17] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt         *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                // 更新时间//[18] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt         *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                // 删除时间//[19] confirm_time                                   datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ConfirmTime       *model.Time `gorm:"column:confirm_time;type:datetime;" json:"confirm_time"`            // 确认时间
}

var ordersTableInfo = &TableInfo{
	Name: "orders",
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
			Name:               "order_status",
			Comment:            `订单状态0待确认1待发货2已发货3已完成`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrderStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "order_status",
			ProtobufFieldName:  "order_status",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "preferential_price",
			Comment:            `优惠金额`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "PreferentialPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "preferential_price",
			ProtobufFieldName:  "preferential_price",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "goods_name_arr",
			Comment:            `订单商品列表`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "GoodsNameArr",
			GoFieldType:        "string",
			JSONFieldName:      "goods_name_arr",
			ProtobufFieldName:  "goods_name_arr",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "remark",
			Comment:            `买家备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "ispay",
			Comment:            `是否已支付`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Ispay",
			GoFieldType:        "int32",
			JSONFieldName:      "ispay",
			ProtobufFieldName:  "ispay",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "pay_price",
			Comment:            `收取金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "PayPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "pay_price",
			ProtobufFieldName:  "pay_price",
			ProtobufType:       "float",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "status",
			Comment:            `此订单状态1正常0禁用-1删除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "pay_time",
			Comment:            `支付时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "PayTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "pay_time",
			ProtobufFieldName:  "pay_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        14,
		},
		&ColumnInfo{
			Index:              14,
			Name:               "ship_time",
			Comment:            `发货时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ShipTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "ship_time",
			ProtobufFieldName:  "ship_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        15,
		},
		&ColumnInfo{
			Index:              15,
			Name:               "complete_time",
			Comment:            `完成时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CompleteTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "complete_time",
			ProtobufFieldName:  "complete_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        16,
		},
		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},
		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},
		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},
		&ColumnInfo{
			Index:              19,
			Name:               "confirm_time",
			Comment:            `确认时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ConfirmTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "confirm_time",
			ProtobufFieldName:  "confirm_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Orders) TableName() string {
	return "orders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Orders) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Orders) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Orders) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Orders) TableInfo() *TableInfo {
	return ordersTableInfo
}

// GetAllOrders is a function to get a slice of record(s) from orders table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (o *Orders) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*Orders, totalRows int64, err error) {

	resultOrm := DB.Model(&Orders{})
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

// GetOrders is a function to get a single record from the orders table in the daigou database
// error - model.ErrNotFound, db Find error
func (o *Orders) Get(DB *gorm.DB, argId uint32) (record *Orders, err error) {
	record = &Orders{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOrders is a function to add a single record to orders table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (o *Orders) Add(DB *gorm.DB, record *Orders) (result *Orders, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOrders is a function to update a single record from orders table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (o *Orders) Update(DB *gorm.DB, argId uint32, updated *Orders) (result *Orders, RowsAffected int64, err error) {

	result = &Orders{}
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

// DeleteOrders is a function to delete a single record from orders table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (o *Orders) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &Orders{}
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
