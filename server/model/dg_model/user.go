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


CREATE TABLE `user` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长用户id',
  `role` int(11) NOT NULL COMMENT '在此平台角色1用户2个体商户3推广员4代购5总代',
  `union_id` varchar(45) COLLATE utf8mb4_bin NOT NULL COMMENT '微信唯一索引',
  `open_id` varchar(45) COLLATE utf8mb4_bin NOT NULL COMMENT '微信openid',
  `gzh_open_id` varchar(45) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '公众号openid',
  `nick_name` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '微信名字',
  `avatar_url` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '微信头像',
  `small_wx_code` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '小程序码',
  `gender` tinyint(4) NOT NULL COMMENT '0未知1男生2女生',
  `country` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '所在国家',
  `province` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '身份',
  `city` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '城市',
  `language` varchar(100) COLLATE utf8mb4_bin NOT NULL COMMENT '语种',
  `bind_shop_id` int(11) DEFAULT NULL COMMENT '绑定的商店id',
  `phone_number` varchar(12) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号码',
  `country_code` varchar(18) COLLATE utf8mb4_bin DEFAULT '86' COMMENT '手机区号',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `open_id` (`open_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000017 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AVG_ROW_LENGTH=1 ROW_FORMAT=DYNAMIC COMMENT='用户表'


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// User struct is a row record of the user table in the daigou database
type User struct {
	//[ 0] user_id                                        uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	UserID      uint32      `gorm:"primary_key;AUTO_INCREMENT;column:user_id;type:uint;" json:"user_id"` // 自增长用户id//[ 1] role                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Role        int32       `gorm:"column:role;type:int;" json:"role"`                                   // 在此平台角色1用户2个体商户3推广员4代购5总代//[ 2] union_id                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	UnionID     string      `gorm:"column:union_id;type:varchar;" json:"union_id"`                       // 微信唯一索引//[ 3] open_id                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	OpenID      string      `gorm:"column:open_id;type:varchar;" json:"open_id"`                         // 微信openid//[ 4] gzh_open_id                                    varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	GzhOpenID   string      `gorm:"column:gzh_open_id;type:varchar;" json:"gzh_open_id"`                 // 公众号openid//[ 5] nick_name                                      varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	NickName    string      `gorm:"column:nick_name;type:varchar;" json:"nick_name"`                     // 微信名字//[ 6] avatar_url                                     varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	AvatarURL   string      `gorm:"column:avatar_url;type:varchar;" json:"avatar_url"`                   // 微信头像//[ 7] small_wx_code                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	SmallWxCode string      `gorm:"column:small_wx_code;type:varchar;" json:"small_wx_code"`             // 小程序码//[ 8] gender                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	Gender      int32       `gorm:"column:gender;type:tinyint;" json:"gender"`                           // 0未知1男生2女生//[ 9] country                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Country     string      `gorm:"column:country;type:varchar;" json:"country"`                         // 所在国家//[10] province                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Province    string      `gorm:"column:province;type:varchar;" json:"province"`                       // 身份//[11] city                                           varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	City        string      `gorm:"column:city;type:varchar;" json:"city"`                               // 城市//[12] language                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Language    string      `gorm:"column:language;type:varchar;" json:"language"`                       // 语种//[13] bind_shop_id                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BindShopID  int32       `gorm:"column:bind_shop_id;type:int;" json:"bind_shop_id"`                   // 绑定的商店id//[14] phone_number                                   varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	PhoneNumber string      `gorm:"column:phone_number;type:varchar;" json:"phone_number"`               // 手机号码//[15] country_code                                   varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: [86]
	CountryCode string      `gorm:"column:country_code;type:varchar;default:86;" json:"country_code"`    // 手机区号//[16] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt   *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                  // 注册时间//[17] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt   *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                  // 修改时间//[18] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt   *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                  // 删除时间
}

var userTableInfo = &TableInfo{
	Name: "user",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "user_id",
			Comment:            `自增长用户id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},
		&ColumnInfo{
			Index:              1,
			Name:               "role",
			Comment:            `在此平台角色1用户2个体商户3推广员4代购5总代`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Role",
			GoFieldType:        "int32",
			JSONFieldName:      "role",
			ProtobufFieldName:  "role",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "union_id",
			Comment:            `微信唯一索引`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "UnionID",
			GoFieldType:        "string",
			JSONFieldName:      "union_id",
			ProtobufFieldName:  "union_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "open_id",
			Comment:            `微信openid`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "OpenID",
			GoFieldType:        "string",
			JSONFieldName:      "open_id",
			ProtobufFieldName:  "open_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "gzh_open_id",
			Comment:            `公众号openid`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "GzhOpenID",
			GoFieldType:        "string",
			JSONFieldName:      "gzh_open_id",
			ProtobufFieldName:  "gzh_open_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "nick_name",
			Comment:            `微信名字`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "NickName",
			GoFieldType:        "string",
			JSONFieldName:      "nick_name",
			ProtobufFieldName:  "nick_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "avatar_url",
			Comment:            `微信头像`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "AvatarURL",
			GoFieldType:        "string",
			JSONFieldName:      "avatar_url",
			ProtobufFieldName:  "avatar_url",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "small_wx_code",
			Comment:            `小程序码`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "SmallWxCode",
			GoFieldType:        "string",
			JSONFieldName:      "small_wx_code",
			ProtobufFieldName:  "small_wx_code",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "gender",
			Comment:            `0未知1男生2女生`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Gender",
			GoFieldType:        "int32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "country",
			Comment:            `所在国家`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Country",
			GoFieldType:        "string",
			JSONFieldName:      "country",
			ProtobufFieldName:  "country",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "province",
			Comment:            `身份`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Province",
			GoFieldType:        "string",
			JSONFieldName:      "province",
			ProtobufFieldName:  "province",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "city",
			Comment:            `城市`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "City",
			GoFieldType:        "string",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "language",
			Comment:            `语种`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Language",
			GoFieldType:        "string",
			JSONFieldName:      "language",
			ProtobufFieldName:  "language",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "bind_shop_id",
			Comment:            `绑定的商店id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BindShopID",
			GoFieldType:        "int32",
			JSONFieldName:      "bind_shop_id",
			ProtobufFieldName:  "bind_shop_id",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
		&ColumnInfo{
			Index:              14,
			Name:               "phone_number",
			Comment:            `手机号码`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "PhoneNumber",
			GoFieldType:        "string",
			JSONFieldName:      "phone_number",
			ProtobufFieldName:  "phone_number",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},
		&ColumnInfo{
			Index:              15,
			Name:               "country_code",
			Comment:            `手机区号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "CountryCode",
			GoFieldType:        "string",
			JSONFieldName:      "country_code",
			ProtobufFieldName:  "country_code",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
		&ColumnInfo{
			Index:              16,
			Name:               "created_at",
			Comment:            `注册时间`,
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
			ProtobufPos:        17,
		},
		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},
		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *User) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *User) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *User) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *User) TableInfo() *TableInfo {
	return userTableInfo
}

// GetAllUser is a function to get a slice of record(s) from user table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (u *User) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*User, totalRows int64, err error) {

	resultOrm := DB.Model(&User{})
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

// GetUser is a function to get a single record from the user table in the daigou database
// error - model.ErrNotFound, db Find error
func (u *User) Get(DB *gorm.DB, argUserID uint32) (record *User, err error) {
	record = &User{}
	if err = DB.First(record, argUserID).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUser is a function to add a single record to user table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (u *User) Add(DB *gorm.DB, record *User) (result *User, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUser is a function to update a single record from user table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (u *User) Update(DB *gorm.DB, argUserID uint32, updated *User) (result *User, RowsAffected int64, err error) {

	result = &User{}
	db := DB.First(result, argUserID)
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

// DeleteUser is a function to delete a single record from user table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (u *User) Delete(DB *gorm.DB, argUserID uint32) (rowsAffected int64, err error) {

	record := &User{}
	db := DB.First(record, argUserID)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
