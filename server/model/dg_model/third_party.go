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


CREATE TABLE `third_party` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `third_company` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `third_account_id` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `third_account_key` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC



*/

// ThirdParty struct is a row record of the third_party table in the daigou database
type ThirdParty struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID              int32       `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`        //[ 1] third_company                                  varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ThirdCompany    string      `gorm:"column:third_company;type:varchar;" json:"third_company"`         //[ 2] third_account_id                               varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ThirdAccountID  string      `gorm:"column:third_account_id;type:varchar;" json:"third_account_id"`   //[ 3] third_account_key                              varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	ThirdAccountKey string      `gorm:"column:third_account_key;type:varchar;" json:"third_account_key"` //[ 4] create_at                                      datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreateAt        *model.Time `gorm:"column:create_at;type:datetime;" json:"create_at"`
}

var third_partyTableInfo = &TableInfo{
	Name: "third_party",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},
		&ColumnInfo{
			Index:              1,
			Name:               "third_company",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ThirdCompany",
			GoFieldType:        "string",
			JSONFieldName:      "third_company",
			ProtobufFieldName:  "third_company",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "third_account_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ThirdAccountID",
			GoFieldType:        "string",
			JSONFieldName:      "third_account_id",
			ProtobufFieldName:  "third_account_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "third_account_key",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "ThirdAccountKey",
			GoFieldType:        "string",
			JSONFieldName:      "third_account_key",
			ProtobufFieldName:  "third_account_key",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "create_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreateAt",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "create_at",
			ProtobufFieldName:  "create_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *ThirdParty) TableName() string {
	return "third_party"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *ThirdParty) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *ThirdParty) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *ThirdParty) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *ThirdParty) TableInfo() *TableInfo {
	return third_partyTableInfo
}

// GetAllThirdParty is a function to get a slice of record(s) from third_party table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (t *ThirdParty) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*ThirdParty, totalRows int64, err error) {

	resultOrm := DB.Model(&ThirdParty{})
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

// GetThirdParty is a function to get a single record from the third_party table in the daigou database
// error - model.ErrNotFound, db Find error
func (t *ThirdParty) Get(DB *gorm.DB, argId int32) (record *ThirdParty, err error) {
	record = &ThirdParty{}
	if err = DB.First(record, argId).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddThirdParty is a function to add a single record to third_party table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (t *ThirdParty) Add(DB *gorm.DB, record *ThirdParty) (result *ThirdParty, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateThirdParty is a function to update a single record from third_party table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (t *ThirdParty) Update(DB *gorm.DB, argId int32, updated *ThirdParty) (result *ThirdParty, RowsAffected int64, err error) {

	result = &ThirdParty{}
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

// DeleteThirdParty is a function to delete a single record from third_party table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (t *ThirdParty) Delete(DB *gorm.DB, argId int32) (rowsAffected int64, err error) {

	record := &ThirdParty{}
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
