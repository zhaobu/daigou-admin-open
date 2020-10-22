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


CREATE TABLE `user_address` (
  `address_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '地址编号',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `phone_number` bigint(20) NOT NULL COMMENT '手机号',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '姓名',
  `gender` int(11) NOT NULL DEFAULT '0' COMMENT '性别,0男生1女生',
  `province` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '省',
  `city` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '市',
  `area` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '区',
  `detailed_address` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '详细地址',
  `is_default` int(11) NOT NULL DEFAULT '0' COMMENT '0:不是默认地址 1:是',
  `classification` int(11) NOT NULL COMMENT '0:代购自己地址 1;名下成员员地址',
  PRIMARY KEY (`address_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 4] column is set for unsigned



*/

// UserAddress struct is a row record of the user_address table in the daigou database
type UserAddress struct {
	//[ 0] address_id                                     uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	AddressID       uint32      `gorm:"primary_key;AUTO_INCREMENT;column:address_id;type:uint;" json:"address_id"` // 地址编号//[ 1] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt       *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                        // 创建时间//[ 2] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt       *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                        // 更新时间//[ 3] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt       *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                        // 删除时间//[ 4] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID          uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                                  // 用户id//[ 5] phone_number                                   bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PhoneNumber     int64       `gorm:"column:phone_number;type:bigint;" json:"phone_number"`                      // 手机号//[ 6] name                                           varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Name            string      `gorm:"column:name;type:varchar;" json:"name"`                                     // 姓名//[ 7] gender                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Gender          int32       `gorm:"column:gender;type:int;default:0;" json:"gender"`                           // 性别,0男生1女生//[ 8] province                                       varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Province        string      `gorm:"column:province;type:varchar;" json:"province"`                             // 省//[ 9] city                                           varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	City            string      `gorm:"column:city;type:varchar;" json:"city"`                                     // 市//[10] area                                           varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Area            string      `gorm:"column:area;type:varchar;" json:"area"`                                     // 区//[11] detailed_address                               varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	DetailedAddress string      `gorm:"column:detailed_address;type:varchar;" json:"detailed_address"`             // 详细地址//[12] is_default                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsDefault       int32       `gorm:"column:is_default;type:int;default:0;" json:"is_default"`                   // 0:不是默认地址 1:是//[13] classification                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Classification  int32       `gorm:"column:classification;type:int;" json:"classification"`                     // 0:代购自己地址 1;名下成员员地址
}

var user_addressTableInfo = &TableInfo{
	Name: "user_address",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "address_id",
			Comment:            `地址编号`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AddressID",
			GoFieldType:        "uint32",
			JSONFieldName:      "address_id",
			ProtobufFieldName:  "address_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},
		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "phone_number",
			Comment:            `手机号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PhoneNumber",
			GoFieldType:        "int64",
			JSONFieldName:      "phone_number",
			ProtobufFieldName:  "phone_number",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "name",
			Comment:            `姓名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "gender",
			Comment:            `性别,0男生1女生`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Gender",
			GoFieldType:        "int32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "province",
			Comment:            `省`,
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
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "city",
			Comment:            `市`,
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
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "area",
			Comment:            `区`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Area",
			GoFieldType:        "string",
			JSONFieldName:      "area",
			ProtobufFieldName:  "area",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "detailed_address",
			Comment:            `详细地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "DetailedAddress",
			GoFieldType:        "string",
			JSONFieldName:      "detailed_address",
			ProtobufFieldName:  "detailed_address",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "is_default",
			Comment:            `0:不是默认地址 1:是`,
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
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "classification",
			Comment:            `0:代购自己地址 1;名下成员员地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Classification",
			GoFieldType:        "int32",
			JSONFieldName:      "classification",
			ProtobufFieldName:  "classification",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserAddress) TableName() string {
	return "user_address"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserAddress) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserAddress) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserAddress) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserAddress) TableInfo() *TableInfo {
	return user_addressTableInfo
}

// GetAllUserAddress is a function to get a slice of record(s) from user_address table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (u *UserAddress) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*UserAddress, totalRows int64, err error) {

	resultOrm := DB.Model(&UserAddress{})
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

// GetUserAddress is a function to get a single record from the user_address table in the daigou database
// error - model.ErrNotFound, db Find error
func (u *UserAddress) Get(DB *gorm.DB, argAddressID uint32) (record *UserAddress, err error) {
	record = &UserAddress{}
	if err = DB.First(record, argAddressID).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserAddress is a function to add a single record to user_address table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (u *UserAddress) Add(DB *gorm.DB, record *UserAddress) (result *UserAddress, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserAddress is a function to update a single record from user_address table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (u *UserAddress) Update(DB *gorm.DB, argAddressID uint32, updated *UserAddress) (result *UserAddress, RowsAffected int64, err error) {

	result = &UserAddress{}
	db := DB.First(result, argAddressID)
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

// DeleteUserAddress is a function to delete a single record from user_address table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (u *UserAddress) Delete(DB *gorm.DB, argAddressID uint32) (rowsAffected int64, err error) {

	record := &UserAddress{}
	db := DB.First(record, argAddressID)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
