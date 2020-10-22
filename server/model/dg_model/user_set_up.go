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


CREATE TABLE `user_set_up` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `privacy_policy` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '隐私政策',
  `user_service_agreement` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户服务协议',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// UserSetUp struct is a row record of the user_set_up table in the daigou database
type UserSetUp struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID                   uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`                 //[ 1] privacy_policy                                 varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	PrivacyPolicy        string `gorm:"column:privacy_policy;type:varchar;" json:"privacy_policy"`                 // 隐私政策//[ 2] user_service_agreement                         varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	UserServiceAgreement string `gorm:"column:user_service_agreement;type:varchar;" json:"user_service_agreement"` // 用户服务协议
}

var user_set_upTableInfo = &TableInfo{
	Name: "user_set_up",
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
			Name:               "privacy_policy",
			Comment:            `隐私政策`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "PrivacyPolicy",
			GoFieldType:        "string",
			JSONFieldName:      "privacy_policy",
			ProtobufFieldName:  "privacy_policy",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "user_service_agreement",
			Comment:            `用户服务协议`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "UserServiceAgreement",
			GoFieldType:        "string",
			JSONFieldName:      "user_service_agreement",
			ProtobufFieldName:  "user_service_agreement",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UserSetUp) TableName() string {
	return "user_set_up"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserSetUp) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserSetUp) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserSetUp) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UserSetUp) TableInfo() *TableInfo {
	return user_set_upTableInfo
}

// GetAllUserSetUp is a function to get a slice of record(s) from user_set_up table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (u *UserSetUp) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*UserSetUp, totalRows int64, err error) {

	resultOrm := DB.Model(&UserSetUp{})
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

// GetUserSetUp is a function to get a single record from the user_set_up table in the daigou database
// error - model.ErrNotFound, db Find error
func (u *UserSetUp) Get(DB *gorm.DB, argId uint32) (record *UserSetUp, err error) {
	record = &UserSetUp{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddUserSetUp is a function to add a single record to user_set_up table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (u *UserSetUp) Add(DB *gorm.DB, record *UserSetUp) (result *UserSetUp, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateUserSetUp is a function to update a single record from user_set_up table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (u *UserSetUp) Update(DB *gorm.DB, argId uint32, updated *UserSetUp) (result *UserSetUp, RowsAffected int64, err error) {

	result = &UserSetUp{}
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

// DeleteUserSetUp is a function to delete a single record from user_set_up table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (u *UserSetUp) Delete(DB *gorm.DB, argId uint32) (rowsAffected int64, err error) {

	record := &UserSetUp{}
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
