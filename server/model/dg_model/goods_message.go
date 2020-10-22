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


CREATE TABLE `goods_message` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `goods_id` bigint(20) unsigned NOT NULL COMMENT '商品id',
  `userid` int(11) DEFAULT NULL COMMENT '留言用户',
  `comment` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '留言内容',
  `created_at` datetime DEFAULT NULL COMMENT '留言时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='商品留言'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// GoodsMessage struct is a row record of the goods_message table in the daigou database
type GoodsMessage struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID        uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] goods_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	GoodsID   uint64      `gorm:"column:goods_id;type:ubigint;" json:"goods_id"`             // 商品id//[ 2] userid                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Userid    int32       `gorm:"column:userid;type:int;" json:"userid"`                     // 留言用户//[ 3] comment                                        varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Comment   string      `gorm:"column:comment;type:varchar;" json:"comment"`               // 留言内容//[ 4] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        // 留言时间
}

var goods_messageTableInfo = &TableInfo{
	Name: "goods_message",
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
			Name:               "goods_id",
			Comment:            `商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "GoodsID",
			GoFieldType:        "uint64",
			JSONFieldName:      "goods_id",
			ProtobufFieldName:  "goods_id",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "userid",
			Comment:            `留言用户`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Userid",
			GoFieldType:        "int32",
			JSONFieldName:      "userid",
			ProtobufFieldName:  "userid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "comment",
			Comment:            `留言内容`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Comment",
			GoFieldType:        "string",
			JSONFieldName:      "comment",
			ProtobufFieldName:  "comment",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "created_at",
			Comment:            `留言时间`,
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
	},
}

// TableName sets the insert table name for this struct type
func (g *GoodsMessage) TableName() string {
	return "goods_message"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GoodsMessage) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GoodsMessage) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GoodsMessage) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GoodsMessage) TableInfo() *TableInfo {
	return goods_messageTableInfo
}

// GetAllGoodsMessage is a function to get a slice of record(s) from goods_message table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (g *GoodsMessage) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*GoodsMessage, totalRows int64, err error) {

	resultOrm := DB.Model(&GoodsMessage{})
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

// GetGoodsMessage is a function to get a single record from the goods_message table in the daigou database
// error - model.ErrNotFound, db Find error
func (g *GoodsMessage) Get(DB *gorm.DB, argId uint32) (record *GoodsMessage, err error) {
	record = &GoodsMessage{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddGoodsMessage is a function to add a single record to goods_message table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (g *GoodsMessage) Add(DB *gorm.DB, record *GoodsMessage) (result *GoodsMessage, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGoodsMessage is a function to update a single record from goods_message table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (g *GoodsMessage) Update(DB *gorm.DB, argId uint32, updated *GoodsMessage) (result *GoodsMessage, RowsAffected int64, err error) {

	result = &GoodsMessage{}
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

// DeleteGoodsMessage is a function to delete a single record from goods_message table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (g *GoodsMessage) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &GoodsMessage{}
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
