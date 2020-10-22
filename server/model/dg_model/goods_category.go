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


CREATE TABLE `goods_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `category_id` int(11) NOT NULL COMMENT '分类id',
  `category_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '分类名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='商品分类表'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// GoodsCategory struct is a row record of the goods_category table in the daigou database
type GoodsCategory struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID           uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] category_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CategoryID   int32  `gorm:"column:category_id;type:int;" json:"category_id"`           // 分类id//[ 2] category_name                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	CategoryName string `gorm:"column:category_name;type:varchar;" json:"category_name"`   // 分类名称
}

var goods_categoryTableInfo = &TableInfo{
	Name: "goods_category",
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
			Name:               "category_id",
			Comment:            `分类id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "int32",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GoodsCategory) TableName() string {
	return "goods_category"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GoodsCategory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GoodsCategory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GoodsCategory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GoodsCategory) TableInfo() *TableInfo {
	return goods_categoryTableInfo
}

// GetAllGoodsCategory is a function to get a slice of record(s) from goods_category table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (g *GoodsCategory) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*GoodsCategory, totalRows int64, err error) {

	resultOrm := DB.Model(&GoodsCategory{})
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

// GetGoodsCategory is a function to get a single record from the goods_category table in the daigou database
// error - model.ErrNotFound, db Find error
func (g *GoodsCategory) Get(DB *gorm.DB, argId uint32) (record *GoodsCategory, err error) {
	record = &GoodsCategory{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddGoodsCategory is a function to add a single record to goods_category table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (g *GoodsCategory) Add(DB *gorm.DB, record *GoodsCategory) (result *GoodsCategory, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGoodsCategory is a function to update a single record from goods_category table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (g *GoodsCategory) Update(DB *gorm.DB, argId uint32, updated *GoodsCategory) (result *GoodsCategory, RowsAffected int64, err error) {

	result = &GoodsCategory{}
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

// DeleteGoodsCategory is a function to delete a single record from goods_category table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (g *GoodsCategory) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &GoodsCategory{}
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
