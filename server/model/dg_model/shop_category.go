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


CREATE TABLE `shop_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `category_id` int(10) unsigned NOT NULL COMMENT '商品分类id',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '店铺ID',
  `category_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '分类名称',
  `category_sort` int(10) unsigned NOT NULL COMMENT '排序字段',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 4] column is set for unsigned



*/

// ShopCategory struct is a row record of the shop_category table in the daigou database
type ShopCategory struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID           uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // id//[ 1] category_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	CategoryID   uint32      `gorm:"column:category_id;type:uint;" json:"category_id"`          // 商品分类id//[ 2] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID       uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`               // 店铺ID//[ 3] category_name                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	CategoryName string      `gorm:"column:category_name;type:varchar;" json:"category_name"`   // 分类名称//[ 4] category_sort                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	CategorySort uint32      `gorm:"column:category_sort;type:uint;" json:"category_sort"`      // 排序字段//[ 5] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt    *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 创建时间
}

var shop_categoryTableInfo = &TableInfo{
	Name: "shop_category",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
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
			Name:               "category_id",
			Comment:            `商品分类id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "uint32",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "shop_id",
			Comment:            `店铺ID`,
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
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "category_name",
			Comment:            `分类名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "CategoryName",
			GoFieldType:        "string",
			JSONFieldName:      "category_name",
			ProtobufFieldName:  "category_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "category_sort",
			Comment:            `排序字段`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CategorySort",
			GoFieldType:        "uint32",
			JSONFieldName:      "category_sort",
			ProtobufFieldName:  "category_sort",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopCategory) TableName() string {
	return "shop_category"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopCategory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopCategory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopCategory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopCategory) TableInfo() *TableInfo {
	return shop_categoryTableInfo
}

// GetAllShopCategory is a function to get a slice of record(s) from shop_category table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopCategory) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopCategory, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopCategory{})
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

// GetShopCategory is a function to get a single record from the shop_category table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopCategory) Get(DB *gorm.DB, argId uint32) (record *ShopCategory, err error) {
	record = &ShopCategory{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopCategory is a function to add a single record to shop_category table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopCategory) Add(DB *gorm.DB, record *ShopCategory) (result *ShopCategory, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopCategory is a function to update a single record from shop_category table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopCategory) Update(DB *gorm.DB, argId uint32, updated *ShopCategory) (result *ShopCategory, RowsAffected int64, err error) {

	result = &ShopCategory{}
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

// DeleteShopCategory is a function to delete a single record from shop_category table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopCategory) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopCategory{}
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
