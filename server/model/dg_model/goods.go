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


CREATE TABLE `goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `goods_id` bigint(20) unsigned NOT NULL COMMENT '商品id',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '商品所属商铺id',
  `goods_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '商品名字',
  `goods_comment` varchar(1000) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品说明',
  `goods_source` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品来源',
  `goods_img_url` json DEFAULT NULL COMMENT '商品图片（oss链接，一次性加载多张图片）',
  `goods_status` int(11) NOT NULL COMMENT '商品状态0下架1上架2售空3即将过期4已过期5删除',
  `category_id` int(11) NOT NULL COMMENT '分类编号',
  `price` decimal(20,2) NOT NULL COMMENT '商品价格',
  `specifications` json NOT NULL COMMENT '规格信息（以Json格式保存）',
  `over_info` text COLLATE utf8mb4_bin COMMENT '过期信息',
  `input_time` datetime DEFAULT NULL COMMENT '入库时间',
  `produced_time` datetime DEFAULT NULL COMMENT '生产日期',
  `add_time` datetime DEFAULT NULL COMMENT '上架时间',
  `down_time` datetime DEFAULT NULL COMMENT '下架时间',
  `top_time` datetime DEFAULT NULL COMMENT '置顶时间',
  `es_de_time` datetime DEFAULT NULL COMMENT '预计发货时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `goodid` (`goods_id`) USING BTREE,
  KEY `shopid` (`shop_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='商品表'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// Goods struct is a row record of the goods table in the daigou database
type Goods struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID             uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] goods_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	GoodsID        uint64      `gorm:"column:goods_id;type:ubigint;" json:"goods_id"`             // 商品id//[ 2] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID         uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`               // 商品所属商铺id//[ 3] goods_name                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	GoodsName      string      `gorm:"column:goods_name;type:varchar;" json:"goods_name"`         // 商品名字//[ 4] goods_comment                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	GoodsComment   string      `gorm:"column:goods_comment;type:varchar;" json:"goods_comment"`   // 商品说明//[ 5] goods_source                                   varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	GoodsSource    string      `gorm:"column:goods_source;type:varchar;" json:"goods_source"`     // 商品来源//[ 6] goods_img_url                                  json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	GoodsImgURL    string      `gorm:"column:goods_img_url;type:json;" json:"goods_img_url"`      // 商品图片（oss链接，一次性加载多张图片）//[ 7] goods_status                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	GoodsStatus    int32       `gorm:"column:goods_status;type:int;" json:"goods_status"`         // 商品状态0下架1上架2售空3即将过期4已过期5删除//[ 8] category_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CategoryID     int32       `gorm:"column:category_id;type:int;" json:"category_id"`           // 分类编号//[ 9] price                                          decimal              null: false  primary: false  isArray: false  auto: false  col: decimal         len: -1      default: []
	Price          float64     `gorm:"column:price;type:decimal;" json:"price"`                   // 商品价格//[10] specifications                                 json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Specifications string      `gorm:"column:specifications;type:json;" json:"specifications"`    // 规格信息（以Json格式保存）//[11] over_info                                      text                 null: true   primary: false  isArray: false  auto: false  col: text            len: 0       default: []
	OverInfo       string      `gorm:"column:over_info;type:text;" json:"over_info"`              // 过期信息//[12] input_time                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	InputTime      *model.Time `gorm:"column:input_time;type:datetime;" json:"input_time"`        // 入库时间//[13] produced_time                                  datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ProducedTime   *model.Time `gorm:"column:produced_time;type:datetime;" json:"produced_time"`  // 生产日期//[14] add_time                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	AddTime        *model.Time `gorm:"column:add_time;type:datetime;" json:"add_time"`            // 上架时间//[15] down_time                                      datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DownTime       *model.Time `gorm:"column:down_time;type:datetime;" json:"down_time"`          // 下架时间//[16] top_time                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	TopTime        *model.Time `gorm:"column:top_time;type:datetime;" json:"top_time"`            // 置顶时间//[17] es_de_time                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	EsDeTime       *model.Time `gorm:"column:es_de_time;type:datetime;" json:"es_de_time"`        // 预计发货时间
}

var goodsTableInfo = &TableInfo{
	Name: "goods",
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
			Name:               "shop_id",
			Comment:            `商品所属商铺id`,
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
			Name:               "goods_name",
			Comment:            `商品名字`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "goods_comment",
			Comment:            `商品说明`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "GoodsComment",
			GoFieldType:        "string",
			JSONFieldName:      "goods_comment",
			ProtobufFieldName:  "goods_comment",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "goods_source",
			Comment:            `商品来源`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "GoodsSource",
			GoFieldType:        "string",
			JSONFieldName:      "goods_source",
			ProtobufFieldName:  "goods_source",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "goods_img_url",
			Comment:            `商品图片（oss链接，一次性加载多张图片）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "GoodsImgURL",
			GoFieldType:        "string",
			JSONFieldName:      "goods_img_url",
			ProtobufFieldName:  "goods_img_url",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "goods_status",
			Comment:            `商品状态0下架1上架2售空3即将过期4已过期5删除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GoodsStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "goods_status",
			ProtobufFieldName:  "goods_status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "category_id",
			Comment:            `分类编号`,
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
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "price",
			Comment:            `商品价格`,
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
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "specifications",
			Comment:            `规格信息（以Json格式保存）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "Specifications",
			GoFieldType:        "string",
			JSONFieldName:      "specifications",
			ProtobufFieldName:  "specifications",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "over_info",
			Comment:            `过期信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       0,
			GoFieldName:        "OverInfo",
			GoFieldType:        "string",
			JSONFieldName:      "over_info",
			ProtobufFieldName:  "over_info",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "input_time",
			Comment:            `入库时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "InputTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "input_time",
			ProtobufFieldName:  "input_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "produced_time",
			Comment:            `生产日期`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ProducedTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "produced_time",
			ProtobufFieldName:  "produced_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        14,
		},
		&ColumnInfo{
			Index:              14,
			Name:               "add_time",
			Comment:            `上架时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        15,
		},
		&ColumnInfo{
			Index:              15,
			Name:               "down_time",
			Comment:            `下架时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "DownTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "down_time",
			ProtobufFieldName:  "down_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        16,
		},
		&ColumnInfo{
			Index:              16,
			Name:               "top_time",
			Comment:            `置顶时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "TopTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "top_time",
			ProtobufFieldName:  "top_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        17,
		},
		&ColumnInfo{
			Index:              17,
			Name:               "es_de_time",
			Comment:            `预计发货时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "EsDeTime",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "es_de_time",
			ProtobufFieldName:  "es_de_time",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        18,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *Goods) TableName() string {
	return "goods"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *Goods) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *Goods) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *Goods) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *Goods) TableInfo() *TableInfo {
	return goodsTableInfo
}

// GetAllGoods is a function to get a slice of record(s) from goods table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (g *Goods) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*Goods, totalRows int64, err error) {

	resultOrm := DB.Model(&Goods{})
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

// GetGoods is a function to get a single record from the goods table in the daigou database
// error - model.ErrNotFound, db Find error
func (g *Goods) Get(DB *gorm.DB, argId uint32) (record *Goods, err error) {
	record = &Goods{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddGoods is a function to add a single record to goods table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (g *Goods) Add(DB *gorm.DB, record *Goods) (result *Goods, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateGoods is a function to update a single record from goods table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (g *Goods) Update(DB *gorm.DB, argId uint32, updated *Goods) (result *Goods, RowsAffected int64, err error) {

	result = &Goods{}
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

// DeleteGoods is a function to delete a single record from goods table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (g *Goods) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &Goods{}
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
