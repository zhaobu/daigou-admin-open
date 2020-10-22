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


CREATE TABLE `shop_express` (
  `express_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '快递标识',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '商铺id',
  `express_code` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '快递编码',
  `express_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '快递名称',
  `express_company_id` int(11) unsigned NOT NULL COMMENT '快递公司id',
  `partner_id` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '快递电子面单账号',
  `partner_key` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '快递电子面子key',
  `express_offer` decimal(10,2) DEFAULT NULL COMMENT '快递报价',
  `express_cost` decimal(10,2) DEFAULT NULL COMMENT '快递成本',
  `is_default` int(11) NOT NULL COMMENT '0不是默认快递1默认快递',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `express_outlets` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '快递网点',
  `temp_id` int(10) unsigned DEFAULT NULL COMMENT '默认模板id',
  PRIMARY KEY (`express_id`) USING BTREE,
  UNIQUE KEY `express_name` (`shop_id`,`express_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 4] column is set for unsigned
[13] column is set for unsigned



*/

// ShopExpress struct is a row record of the shop_express table in the daigou database
type ShopExpress struct {
	//[ 0] express_id                                     uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ExpressID        uint32      `gorm:"primary_key;AUTO_INCREMENT;column:express_id;type:uint;" json:"express_id"` // 快递标识//[ 1] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID           uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`                               // 商铺id//[ 2] express_code                                   varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ExpressCode      string      `gorm:"column:express_code;type:varchar;" json:"express_code"`                     // 快递编码//[ 3] express_name                                   varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ExpressName      string      `gorm:"column:express_name;type:varchar;" json:"express_name"`                     // 快递名称//[ 4] express_company_id                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ExpressCompanyID uint32      `gorm:"column:express_company_id;type:uint;" json:"express_company_id"`            // 快递公司id//[ 5] partner_id                                     varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	PartnerID        string      `gorm:"column:partner_id;type:varchar;" json:"partner_id"`                         // 快递电子面单账号//[ 6] partner_key                                    varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	PartnerKey       string      `gorm:"column:partner_key;type:varchar;" json:"partner_key"`                       // 快递电子面子key//[ 7] express_offer                                  decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	ExpressOffer     float64     `gorm:"column:express_offer;type:decimal;" json:"express_offer"`                   // 快递报价//[ 8] express_cost                                   decimal              null: true   primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	ExpressCost      float64     `gorm:"column:express_cost;type:decimal;" json:"express_cost"`                     // 快递成本//[ 9] is_default                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	IsDefault        int32       `gorm:"column:is_default;type:int;" json:"is_default"`                             // 0不是默认快递1默认快递//[10] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt        *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                        // 创建时间//[11] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt        *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                        // 修改时间//[12] express_outlets                                varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ExpressOutlets   string      `gorm:"column:express_outlets;type:varchar;" json:"express_outlets"`               // 快递网点//[13] temp_id                                        uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TempID           uint32      `gorm:"column:temp_id;type:uint;" json:"temp_id"`                                  // 默认模板id
}

var shop_expressTableInfo = &TableInfo{
	Name: "shop_express",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "express_id",
			Comment:            `快递标识`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressID",
			GoFieldType:        "uint32",
			JSONFieldName:      "express_id",
			ProtobufFieldName:  "express_id",
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
			Name:               "express_code",
			Comment:            `快递编码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ExpressCode",
			GoFieldType:        "string",
			JSONFieldName:      "express_code",
			ProtobufFieldName:  "express_code",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "express_name",
			Comment:            `快递名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ExpressName",
			GoFieldType:        "string",
			JSONFieldName:      "express_name",
			ProtobufFieldName:  "express_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "express_company_id",
			Comment:            `快递公司id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressCompanyID",
			GoFieldType:        "uint32",
			JSONFieldName:      "express_company_id",
			ProtobufFieldName:  "express_company_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "partner_id",
			Comment:            `快递电子面单账号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "PartnerID",
			GoFieldType:        "string",
			JSONFieldName:      "partner_id",
			ProtobufFieldName:  "partner_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "partner_key",
			Comment:            `快递电子面子key`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "PartnerKey",
			GoFieldType:        "string",
			JSONFieldName:      "partner_key",
			ProtobufFieldName:  "partner_key",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "express_offer",
			Comment:            `快递报价`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "ExpressOffer",
			GoFieldType:        "float64",
			JSONFieldName:      "express_offer",
			ProtobufFieldName:  "express_offer",
			ProtobufType:       "float",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "express_cost",
			Comment:            `快递成本`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "ExpressCost",
			GoFieldType:        "float64",
			JSONFieldName:      "express_cost",
			ProtobufFieldName:  "express_cost",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "is_default",
			Comment:            `0不是默认快递1默认快递`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsDefault",
			GoFieldType:        "int32",
			JSONFieldName:      "is_default",
			ProtobufFieldName:  "is_default",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "updated_at",
			Comment:            `修改时间`,
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
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "express_outlets",
			Comment:            `快递网点`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ExpressOutlets",
			GoFieldType:        "string",
			JSONFieldName:      "express_outlets",
			ProtobufFieldName:  "express_outlets",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "temp_id",
			Comment:            `默认模板id`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TempID",
			GoFieldType:        "uint32",
			JSONFieldName:      "temp_id",
			ProtobufFieldName:  "temp_id",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopExpress) TableName() string {
	return "shop_express"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopExpress) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopExpress) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopExpress) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopExpress) TableInfo() *TableInfo {
	return shop_expressTableInfo
}

// GetAllShopExpress is a function to get a slice of record(s) from shop_express table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopExpress) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopExpress, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopExpress{})
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

// GetShopExpress is a function to get a single record from the shop_express table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopExpress) Get(DB *gorm.DB, argExpressID uint32) (record *ShopExpress, err error) {
	record = &ShopExpress{}
	if err = DB.First(record, argExpressID).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopExpress is a function to add a single record to shop_express table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopExpress) Add(DB *gorm.DB, record *ShopExpress) (result *ShopExpress, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopExpress is a function to update a single record from shop_express table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopExpress) Update(DB *gorm.DB, argExpressID uint32, updated *ShopExpress) (result *ShopExpress, RowsAffected int64, err error) {

	result = &ShopExpress{}
	db := DB.First(result, argExpressID)
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

// DeleteShopExpress is a function to delete a single record from shop_express table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopExpress) Delete(DB *gorm.DB, argExpressID uint32) (rowsAffected int64, err error) {

	record := &ShopExpress{}
	db := DB.First(record, argExpressID)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
