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


CREATE TABLE `orders_logistics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单编号',
  `status` int(11) NOT NULL COMMENT '0发送中1收货中2完成-1异常',
  `cost` decimal(10,2) DEFAULT NULL COMMENT '运费成本',
  `offer` decimal(10,2) NOT NULL COMMENT '运费报价',
  `receiver_name` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT '收货人',
  `receiver_iphone` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '收货联系方式',
  `receiver_province` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '省',
  `receiver_city` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '市',
  `receiver_district` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '区/县',
  `receiver_address` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '详细地址',
  `sender_name` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '寄货人',
  `sender_iphone` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '寄货联系方式',
  `sender_province` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '寄货省',
  `sender_city` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '寄货市',
  `sender_district` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '寄货区/县',
  `sender_address` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '寄货详细地址',
  `logistics_company` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '物流公司',
  `logistics_com` varchar(16) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '物流公司编码（如：顺丰编码(SF)）',
  `logistics_number` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '物流编号',
  `third_party_id` int(11) DEFAULT NULL COMMENT '物流第三方平台编号',
  `logistics_records` mediumtext COLLATE utf8mb4_bin COMMENT '快递物流记录',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='订单物流信息表'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// OrdersLogistics struct is a row record of the orders_logistics table in the daigou database
type OrdersLogistics struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID               uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`       // 主键//[ 1] order_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OrderID          string      `gorm:"column:order_id;type:varchar;" json:"order_id"`                   // 订单编号//[ 2] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Status           int32       `gorm:"column:status;type:int;" json:"status"`                           // 0发送中1收货中2完成-1异常//[ 3] cost                                           decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Cost             float64     `gorm:"column:cost;type:decimal;" json:"cost"`                           // 运费成本//[ 4] offer                                          decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Offer            float64     `gorm:"column:offer;type:decimal;" json:"offer"`                         // 运费报价//[ 5] receiver_name                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ReceiverName     string      `gorm:"column:receiver_name;type:varchar;" json:"receiver_name"`         // 收货人//[ 6] receiver_iphone                                varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ReceiverIphone   string      `gorm:"column:receiver_iphone;type:varchar;" json:"receiver_iphone"`     // 收货联系方式//[ 7] receiver_province                              varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ReceiverProvince string      `gorm:"column:receiver_province;type:varchar;" json:"receiver_province"` // 省//[ 8] receiver_city                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ReceiverCity     string      `gorm:"column:receiver_city;type:varchar;" json:"receiver_city"`         // 市//[ 9] receiver_district                              varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ReceiverDistrict string      `gorm:"column:receiver_district;type:varchar;" json:"receiver_district"` // 区/县//[10] receiver_address                               varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ReceiverAddress  string      `gorm:"column:receiver_address;type:varchar;" json:"receiver_address"`   // 详细地址//[11] sender_name                                    varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SenderName       string      `gorm:"column:sender_name;type:varchar;" json:"sender_name"`             // 寄货人//[12] sender_iphone                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SenderIphone     string      `gorm:"column:sender_iphone;type:varchar;" json:"sender_iphone"`         // 寄货联系方式//[13] sender_province                                varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SenderProvince   string      `gorm:"column:sender_province;type:varchar;" json:"sender_province"`     // 寄货省//[14] sender_city                                    varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SenderCity       string      `gorm:"column:sender_city;type:varchar;" json:"sender_city"`             // 寄货市//[15] sender_district                                varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SenderDistrict   string      `gorm:"column:sender_district;type:varchar;" json:"sender_district"`     // 寄货区/县//[16] sender_address                                 varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SenderAddress    string      `gorm:"column:sender_address;type:varchar;" json:"sender_address"`       // 寄货详细地址//[17] logistics_company                              varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	LogisticsCompany string      `gorm:"column:logistics_company;type:varchar;" json:"logistics_company"` // 物流公司//[18] logistics_com                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	LogisticsCom     string      `gorm:"column:logistics_com;type:varchar;" json:"logistics_com"`         // 物流公司编码（如：顺丰编码(SF)）//[19] logistics_number                               varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	LogisticsNumber  string      `gorm:"column:logistics_number;type:varchar;" json:"logistics_number"`   // 物流编号//[20] third_party_id                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ThirdPartyID     int32       `gorm:"column:third_party_id;type:int;" json:"third_party_id"`           // 物流第三方平台编号//[21] logistics_records                              text                 null: true   primary: false  isArray: false  auto: false  col: text            len: 0       default: []
	LogisticsRecords string      `gorm:"column:logistics_records;type:text;" json:"logistics_records"`    // 快递物流记录//[22] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt        *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`              // 创建时间//[23] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt        *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`              // 更新时间//[24] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt        *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`              // 删除时间
}

var orders_logisticsTableInfo = &TableInfo{
	Name: "orders_logistics",
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
			Name:               "status",
			Comment:            `0发送中1收货中2完成-1异常`,
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
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "cost",
			Comment:            `运费成本`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "Cost",
			GoFieldType:        "float64",
			JSONFieldName:      "cost",
			ProtobufFieldName:  "cost",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "offer",
			Comment:            `运费报价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "Offer",
			GoFieldType:        "float64",
			JSONFieldName:      "offer",
			ProtobufFieldName:  "offer",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "receiver_name",
			Comment:            `收货人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ReceiverName",
			GoFieldType:        "string",
			JSONFieldName:      "receiver_name",
			ProtobufFieldName:  "receiver_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "receiver_iphone",
			Comment:            `收货联系方式`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ReceiverIphone",
			GoFieldType:        "string",
			JSONFieldName:      "receiver_iphone",
			ProtobufFieldName:  "receiver_iphone",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "receiver_province",
			Comment:            `省`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ReceiverProvince",
			GoFieldType:        "string",
			JSONFieldName:      "receiver_province",
			ProtobufFieldName:  "receiver_province",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "receiver_city",
			Comment:            `市`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ReceiverCity",
			GoFieldType:        "string",
			JSONFieldName:      "receiver_city",
			ProtobufFieldName:  "receiver_city",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "receiver_district",
			Comment:            `区/县`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ReceiverDistrict",
			GoFieldType:        "string",
			JSONFieldName:      "receiver_district",
			ProtobufFieldName:  "receiver_district",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "receiver_address",
			Comment:            `详细地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ReceiverAddress",
			GoFieldType:        "string",
			JSONFieldName:      "receiver_address",
			ProtobufFieldName:  "receiver_address",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "sender_name",
			Comment:            `寄货人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SenderName",
			GoFieldType:        "string",
			JSONFieldName:      "sender_name",
			ProtobufFieldName:  "sender_name",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "sender_iphone",
			Comment:            `寄货联系方式`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SenderIphone",
			GoFieldType:        "string",
			JSONFieldName:      "sender_iphone",
			ProtobufFieldName:  "sender_iphone",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "sender_province",
			Comment:            `寄货省`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SenderProvince",
			GoFieldType:        "string",
			JSONFieldName:      "sender_province",
			ProtobufFieldName:  "sender_province",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},
		&ColumnInfo{
			Index:              14,
			Name:               "sender_city",
			Comment:            `寄货市`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SenderCity",
			GoFieldType:        "string",
			JSONFieldName:      "sender_city",
			ProtobufFieldName:  "sender_city",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},
		&ColumnInfo{
			Index:              15,
			Name:               "sender_district",
			Comment:            `寄货区/县`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SenderDistrict",
			GoFieldType:        "string",
			JSONFieldName:      "sender_district",
			ProtobufFieldName:  "sender_district",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
		&ColumnInfo{
			Index:              16,
			Name:               "sender_address",
			Comment:            `寄货详细地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SenderAddress",
			GoFieldType:        "string",
			JSONFieldName:      "sender_address",
			ProtobufFieldName:  "sender_address",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},
		&ColumnInfo{
			Index:              17,
			Name:               "logistics_company",
			Comment:            `物流公司`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "LogisticsCompany",
			GoFieldType:        "string",
			JSONFieldName:      "logistics_company",
			ProtobufFieldName:  "logistics_company",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},
		&ColumnInfo{
			Index:              18,
			Name:               "logistics_com",
			Comment:            `物流公司编码（如：顺丰编码(SF)）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "LogisticsCom",
			GoFieldType:        "string",
			JSONFieldName:      "logistics_com",
			ProtobufFieldName:  "logistics_com",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},
		&ColumnInfo{
			Index:              19,
			Name:               "logistics_number",
			Comment:            `物流编号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "LogisticsNumber",
			GoFieldType:        "string",
			JSONFieldName:      "logistics_number",
			ProtobufFieldName:  "logistics_number",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},
		&ColumnInfo{
			Index:              20,
			Name:               "third_party_id",
			Comment:            `物流第三方平台编号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ThirdPartyID",
			GoFieldType:        "int32",
			JSONFieldName:      "third_party_id",
			ProtobufFieldName:  "third_party_id",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},
		&ColumnInfo{
			Index:              21,
			Name:               "logistics_records",
			Comment:            `快递物流记录`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       0,
			GoFieldName:        "LogisticsRecords",
			GoFieldType:        "string",
			JSONFieldName:      "logistics_records",
			ProtobufFieldName:  "logistics_records",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},
		&ColumnInfo{
			Index:              22,
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
			ProtobufPos:        23,
		},
		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},
		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrdersLogistics) TableName() string {
	return "orders_logistics"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrdersLogistics) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrdersLogistics) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrdersLogistics) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrdersLogistics) TableInfo() *TableInfo {
	return orders_logisticsTableInfo
}

// GetAllOrdersLogistics is a function to get a slice of record(s) from orders_logistics table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (o *OrdersLogistics) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*OrdersLogistics, totalRows int64, err error) {

	resultOrm := DB.Model(&OrdersLogistics{})
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

// GetOrdersLogistics is a function to get a single record from the orders_logistics table in the daigou database
// error - model.ErrNotFound, db Find error
func (o *OrdersLogistics) Get(DB *gorm.DB, argId uint32) (record *OrdersLogistics, err error) {
	record = &OrdersLogistics{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOrdersLogistics is a function to add a single record to orders_logistics table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (o *OrdersLogistics) Add(DB *gorm.DB, record *OrdersLogistics) (result *OrdersLogistics, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateOrdersLogistics is a function to update a single record from orders_logistics table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (o *OrdersLogistics) Update(DB *gorm.DB, argId uint32, updated *OrdersLogistics) (result *OrdersLogistics, RowsAffected int64, err error) {

	result = &OrdersLogistics{}
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

// DeleteOrdersLogistics is a function to delete a single record from orders_logistics table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (o *OrdersLogistics) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &OrdersLogistics{}
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
