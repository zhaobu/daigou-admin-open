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


CREATE TABLE `shop_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `shop_id` bigint(20) unsigned NOT NULL COMMENT '店铺ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `shop_name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '店铺名称',
  `shop_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '店铺头像',
  `shop_description` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '店铺说明',
  `qr_code_url` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '收款二维码',
  `shop_fans_count` int(11) NOT NULL DEFAULT '0' COMMENT '店铺粉丝数',
  `wechat_number` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '微信号',
  `is_recommend` int(10) unsigned NOT NULL COMMENT '0:不推荐 1:推荐店铺',
  `is_enable` int(11) NOT NULL DEFAULT '1' COMMENT '商铺状态0不启用1启用',
  `category_info` json DEFAULT NULL COMMENT '用户分类信息',
  `mainpage_scroll_info` json DEFAULT NULL COMMENT '用户首页滚动信息',
  `shop_watermark` json DEFAULT NULL COMMENT '图片水印配置',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `shop_id` (`shop_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='商铺'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 5] column is set for unsigned
[12] column is set for unsigned



*/

// ShopInfo struct is a row record of the shop_info table in the daigou database
type ShopInfo struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID                 uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`          // 主键//[ 1] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	ShopID             uint64      `gorm:"column:shop_id;type:ubigint;" json:"shop_id"`                        // 店铺ID//[ 2] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt          *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                 // 创建时间//[ 3] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt          *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                 // 更新时间//[ 4] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt          *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                 // 删除时间//[ 5] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID             uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                           // 用户id//[ 6] shop_name                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ShopName           string      `gorm:"column:shop_name;type:varchar;" json:"shop_name"`                    // 店铺名称//[ 7] shop_url                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ShopURL            string      `gorm:"column:shop_url;type:varchar;" json:"shop_url"`                      // 店铺头像//[ 8] shop_description                               varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ShopDescription    string      `gorm:"column:shop_description;type:varchar;" json:"shop_description"`      // 店铺说明//[ 9] qr_code_url                                    varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	QrCodeURL          string      `gorm:"column:qr_code_url;type:varchar;" json:"qr_code_url"`                // 收款二维码//[10] shop_fans_count                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShopFansCount      int32       `gorm:"column:shop_fans_count;type:int;default:0;" json:"shop_fans_count"`  // 店铺粉丝数//[11] wechat_number                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	WechatNumber       string      `gorm:"column:wechat_number;type:varchar;" json:"wechat_number"`            // 微信号//[12] is_recommend                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	IsRecommend        uint32      `gorm:"column:is_recommend;type:uint;" json:"is_recommend"`                 // 0:不推荐 1:推荐店铺//[13] is_enable                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	IsEnable           int32       `gorm:"column:is_enable;type:int;default:1;" json:"is_enable"`              // 商铺状态0不启用1启用//[14] category_info                                  json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	CategoryInfo       string      `gorm:"column:category_info;type:json;" json:"category_info"`               // 用户分类信息//[15] mainpage_scroll_info                           json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	MainpageScrollInfo string      `gorm:"column:mainpage_scroll_info;type:json;" json:"mainpage_scroll_info"` // 用户首页滚动信息//[16] shop_watermark                                 json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	ShopWatermark      string      `gorm:"column:shop_watermark;type:json;" json:"shop_watermark"`             // 图片水印配置
}

var shop_infoTableInfo = &TableInfo{
	Name: "shop_info",
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
			Comment:            `用户id`,
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
			Name:               "shop_name",
			Comment:            `店铺名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ShopName",
			GoFieldType:        "string",
			JSONFieldName:      "shop_name",
			ProtobufFieldName:  "shop_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "shop_url",
			Comment:            `店铺头像`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ShopURL",
			GoFieldType:        "string",
			JSONFieldName:      "shop_url",
			ProtobufFieldName:  "shop_url",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "shop_description",
			Comment:            `店铺说明`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ShopDescription",
			GoFieldType:        "string",
			JSONFieldName:      "shop_description",
			ProtobufFieldName:  "shop_description",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "qr_code_url",
			Comment:            `收款二维码`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "QrCodeURL",
			GoFieldType:        "string",
			JSONFieldName:      "qr_code_url",
			ProtobufFieldName:  "qr_code_url",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "shop_fans_count",
			Comment:            `店铺粉丝数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ShopFansCount",
			GoFieldType:        "int32",
			JSONFieldName:      "shop_fans_count",
			ProtobufFieldName:  "shop_fans_count",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "wechat_number",
			Comment:            `微信号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "WechatNumber",
			GoFieldType:        "string",
			JSONFieldName:      "wechat_number",
			ProtobufFieldName:  "wechat_number",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "is_recommend",
			Comment:            `0:不推荐 1:推荐店铺`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "IsRecommend",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_recommend",
			ProtobufFieldName:  "is_recommend",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "is_enable",
			Comment:            `商铺状态0不启用1启用`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsEnable",
			GoFieldType:        "int32",
			JSONFieldName:      "is_enable",
			ProtobufFieldName:  "is_enable",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
		&ColumnInfo{
			Index:              14,
			Name:               "category_info",
			Comment:            `用户分类信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "CategoryInfo",
			GoFieldType:        "string",
			JSONFieldName:      "category_info",
			ProtobufFieldName:  "category_info",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},
		&ColumnInfo{
			Index:              15,
			Name:               "mainpage_scroll_info",
			Comment:            `用户首页滚动信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "MainpageScrollInfo",
			GoFieldType:        "string",
			JSONFieldName:      "mainpage_scroll_info",
			ProtobufFieldName:  "mainpage_scroll_info",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
		&ColumnInfo{
			Index:              16,
			Name:               "shop_watermark",
			Comment:            `图片水印配置`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "ShopWatermark",
			GoFieldType:        "string",
			JSONFieldName:      "shop_watermark",
			ProtobufFieldName:  "shop_watermark",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopInfo) TableName() string {
	return "shop_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopInfo) TableInfo() *TableInfo {
	return shop_infoTableInfo
}

// GetAllShopInfo is a function to get a slice of record(s) from shop_info table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (s *ShopInfo) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ShopInfo, totalRows int64, err error) {

	resultOrm := DB.Model(&ShopInfo{})
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

// GetShopInfo is a function to get a single record from the shop_info table in the daigou database
// error - model.ErrNotFound, db Find error
func (s *ShopInfo) Get(DB *gorm.DB, argId uint32) (record *ShopInfo, err error) {
	record = &ShopInfo{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddShopInfo is a function to add a single record to shop_info table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (s *ShopInfo) Add(DB *gorm.DB, record *ShopInfo) (result *ShopInfo, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateShopInfo is a function to update a single record from shop_info table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (s *ShopInfo) Update(DB *gorm.DB, argId uint32, updated *ShopInfo) (result *ShopInfo, RowsAffected int64, err error) {

	result = &ShopInfo{}
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

// DeleteShopInfo is a function to delete a single record from shop_info table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (s *ShopInfo) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &ShopInfo{}
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
