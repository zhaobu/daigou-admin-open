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


CREATE TABLE `orders_bill_flow` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单编号',
  `before` decimal(20,2) DEFAULT NULL COMMENT '之前金额',
  `last` decimal(20,2) DEFAULT NULL COMMENT '之后金额',
  `change_value` decimal(20,2) DEFAULT NULL COMMENT '变化金额',
  `remarks` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='订单账单流水记录表'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// OrdersBillFlow struct is a row record of the orders_bill_flow table in the daigou database
type OrdersBillFlow struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID          uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] order_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OrderID     string      `gorm:"column:order_id;type:varchar;" json:"order_id"`             // 订单编号//[ 2] before                                         decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Before      float64     `gorm:"column:before;type:decimal;" json:"before"`                 // 之前金额//[ 3] last                                           decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Last        float64     `gorm:"column:last;type:decimal;" json:"last"`                     // 之后金额//[ 4] change_value                                   decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	ChangeValue float64     `gorm:"column:change_value;type:decimal;" json:"change_value"`     // 变化金额//[ 5] remarks                                        varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Remarks     string      `gorm:"column:remarks;type:varchar;" json:"remarks"`               // 备注//[ 6] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt   *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 创建时间//[ 7] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt   *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`        // 更新时间//[ 8] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt   *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`        // 删除时间
}

var orders_bill_flowTableInfo = &TableInfo{
	Name: "orders_bill_flow",
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
			Name:               "before",
			Comment:            `之前金额`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "Before",
			GoFieldType:        "float64",
			JSONFieldName:      "before",
			ProtobufFieldName:  "before",
			ProtobufType:       "float",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "last",
			Comment:            `之后金额`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "Last",
			GoFieldType:        "float64",
			JSONFieldName:      "last",
			ProtobufFieldName:  "last",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "change_value",
			Comment:            `变化金额`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "ChangeValue",
			GoFieldType:        "float64",
			JSONFieldName:      "change_value",
			ProtobufFieldName:  "change_value",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "remarks",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Remarks",
			GoFieldType:        "string",
			JSONFieldName:      "remarks",
			ProtobufFieldName:  "remarks",
			ProtobufType:       "string",
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
func (o *OrdersBillFlow) TableName() string {
	return "orders_bill_flow"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrdersBillFlow) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrdersBillFlow) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrdersBillFlow) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrdersBillFlow) TableInfo() *TableInfo {
	return orders_bill_flowTableInfo
}

// GetAllOrdersBillFlow is a function to get a slice of record(s) from orders_bill_flow table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (o *OrdersBillFlow) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*OrdersBillFlow, totalRows int64, err error) {

	resultOrm := DB.Model(&OrdersBillFlow{})
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

// GetOrdersBillFlow is a function to get a single record from the orders_bill_flow table in the daigou database
// error - model.ErrNotFound, db Find error
func (o *OrdersBillFlow) Get(DB *gorm.DB, argId uint32) (record *OrdersBillFlow, err error) {
	record = &OrdersBillFlow{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOrdersBillFlow is a function to add a single record to orders_bill_flow table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (o *OrdersBillFlow) Add(DB *gorm.DB, record *OrdersBillFlow) (result *OrdersBillFlow, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOrdersBillFlow is a function to update a single record from orders_bill_flow table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (o *OrdersBillFlow) Update(DB *gorm.DB, argId uint32, updated *OrdersBillFlow) (result *OrdersBillFlow, RowsAffected int64, err error) {

	result = &OrdersBillFlow{}
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

// DeleteOrdersBillFlow is a function to delete a single record from orders_bill_flow table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (o *OrdersBillFlow) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &OrdersBillFlow{}
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
