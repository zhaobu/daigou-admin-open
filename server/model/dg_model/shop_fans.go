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


CREATE TABLE `shop_fans` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '店铺id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_id` int(10) unsigned NOT NULL COMMENT '粉丝id',
  `category` int(11) NOT NULL COMMENT '用于排序：1置顶：为绑定用户 0：为关注用户',
  `transaction_number` int(11) NOT NULL DEFAULT '0' COMMENT '订单数量',
  `transaction_amount` decimal(10,0) NOT NULL COMMENT '订单总金额',
  `end_time` datetime DEFAULT NULL COMMENT '最后购买时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `shop_id_2` (`shop_id`,`user_id`) USING BTREE COMMENT '粉丝绑定唯一约束',
  KEY `shop_id` (`shop_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='代购粉丝表'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 5] column is set for unsigned



*/

// ShopFans struct is a row record of the shop_fans table in the daigou database
type ShopFans struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID                uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`               // 主键//[ 1] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID            uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`                             // 店铺id//[ 2] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt         *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                      // 创建时间//[ 3] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt         *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                      // 更新时间//[ 4] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt         *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                      // 删除时间//[ 5] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID            uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                                // 粉丝id//[ 6] category                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Category          int32       `gorm:"column:category;type:int;" json:"category"`                               // 用于排序：1置顶：为绑定用户 0：为关注用户//[ 7] transaction_number                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TransactionNumber int32       `gorm:"column:transaction_number;type:int;default:0;" json:"transaction_number"` // 订单数量//[ 8] transaction_amount                             decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	TransactionAmount float64     `gorm:"column:transaction_amount;type:decimal;" json:"transaction_amount"`       // 订单总金额//[ 9] end_time                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	EndTime           *model.Time `gorm:"column:end_time;type:datetime;" json:"end_time"`                          // 最后购买时间
}

var shop_fansTableInfo = &TableInfo{
	Name: "shop_fans",
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
			Comment:            `店铺id`,
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
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "user_id",
			Comment:            `粉丝id`,
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
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "category",
			Comment:            `用于排序：1置顶：为绑定用户 0：为关注用户`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Category",
			GoFieldType:        "int32",
			JSONFieldName:      "category",
			ProtobufFieldName:  "category",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "transaction_number",
			Comment:            `订单数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TransactionNumber",
			GoFieldType:        "int32",
			JSONFieldName:      "transaction_number",
			ProtobufFieldName:  "transaction_number",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "transaction_amount",
			Comment:            `订单总金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "decimal",
			DatabaseTypePretty: "decimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "decimal",
			ColumnLength:       -1,
			GoFieldName:        "TransactionAmount",
			GoFieldType:        "float64",
			JSONFieldName:      "transaction_amount",
			ProtobufFieldName:  "transaction_amount",
			ProtobufType:       "float",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "end_time",
			Comment:            `最后购买时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopFans) TableName() string {
	return "shop_fans"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopFans) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopFans) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopFans) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopFans) TableInfo() *TableInfo {
	return shop_fansTableInfo
}

// GetAllShopFans is a function to get a slice of record(s) from shop_fans table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopFans) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopFans, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopFans{})
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

// GetShopFans is a function to get a single record from the shop_fans table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopFans) Get(DB *gorm.DB, argId uint32) (record *ShopFans, err error) {
	record = &ShopFans{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopFans is a function to add a single record to shop_fans table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopFans) Add(DB *gorm.DB, record *ShopFans) (result *ShopFans, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopFans is a function to update a single record from shop_fans table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopFans) Update(DB *gorm.DB, argId uint32, updated *ShopFans) (result *ShopFans, RowsAffected int64, err error) {

	result = &ShopFans{}
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

// DeleteShopFans is a function to delete a single record from shop_fans table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopFans) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopFans{}
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
