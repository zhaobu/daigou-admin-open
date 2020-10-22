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


CREATE TABLE `shop_vip_explain` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `limit_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '限制名称',
  `limit_explain` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '限制说明',
  `vip_explain` int(11) NOT NULL COMMENT 'vip限制（）',
  `ordinary_explain` int(11) NOT NULL COMMENT '普通用户限制（）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// ShopVipExplain struct is a row record of the shop_vip_explain table in the daigou database
type ShopVipExplain struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID              uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` //[ 1] limit_name                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	LimitName       string `gorm:"column:limit_name;type:varchar;" json:"limit_name"`         // 限制名称//[ 2] limit_explain                                  varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	LimitExplain    string `gorm:"column:limit_explain;type:varchar;" json:"limit_explain"`   // 限制说明//[ 3] vip_explain                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	VipExplain      int32  `gorm:"column:vip_explain;type:int;" json:"vip_explain"`           // vip限制（）//[ 4] ordinary_explain                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	OrdinaryExplain int32  `gorm:"column:ordinary_explain;type:int;" json:"ordinary_explain"` // 普通用户限制（）
}

var shop_vip_explainTableInfo = &TableInfo{
	Name: "shop_vip_explain",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
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
			Name:               "limit_name",
			Comment:            `限制名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "LimitName",
			GoFieldType:        "string",
			JSONFieldName:      "limit_name",
			ProtobufFieldName:  "limit_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "limit_explain",
			Comment:            `限制说明`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "LimitExplain",
			GoFieldType:        "string",
			JSONFieldName:      "limit_explain",
			ProtobufFieldName:  "limit_explain",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "vip_explain",
			Comment:            `vip限制（）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "VipExplain",
			GoFieldType:        "int32",
			JSONFieldName:      "vip_explain",
			ProtobufFieldName:  "vip_explain",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "ordinary_explain",
			Comment:            `普通用户限制（）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrdinaryExplain",
			GoFieldType:        "int32",
			JSONFieldName:      "ordinary_explain",
			ProtobufFieldName:  "ordinary_explain",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopVipExplain) TableName() string {
	return "shop_vip_explain"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopVipExplain) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopVipExplain) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopVipExplain) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopVipExplain) TableInfo() *TableInfo {
	return shop_vip_explainTableInfo
}

// GetAllShopVipExplain is a function to get a slice of record(s) from shop_vip_explain table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopVipExplain) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopVipExplain, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopVipExplain{})
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

// GetShopVipExplain is a function to get a single record from the shop_vip_explain table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopVipExplain) Get(DB *gorm.DB, argId uint32) (record *ShopVipExplain, err error) {
	record = &ShopVipExplain{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopVipExplain is a function to add a single record to shop_vip_explain table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopVipExplain) Add(DB *gorm.DB, record *ShopVipExplain) (result *ShopVipExplain, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopVipExplain is a function to update a single record from shop_vip_explain table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopVipExplain) Update(DB *gorm.DB, argId uint32, updated *ShopVipExplain) (result *ShopVipExplain, RowsAffected int64, err error) {

	result = &ShopVipExplain{}
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

// DeleteShopVipExplain is a function to delete a single record from shop_vip_explain table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopVipExplain) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopVipExplain{}
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
