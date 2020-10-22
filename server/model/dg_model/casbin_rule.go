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


CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] Warning table: casbin_rule does not have a primary key defined, setting col position 1 p_type as primary key
Warning table: casbin_rule primary key column p_type is nullable column, setting it as NOT NULL




*/

// CasbinRule struct is a row record of the casbin_rule table in the daigou database
type CasbinRule struct {
	//[ 0] p_type                                         varchar              null: false  primary: true   isArray: false  auto: false  col: varchar         len: 0       default: []
	PType string `gorm:"primary_key;column:p_type;type:varchar;" json:"p_type"` //[ 1] v0                                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	V0    string `gorm:"column:v0;type:varchar;" json:"v_0"`                    //[ 2] v1                                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	V1    string `gorm:"column:v1;type:varchar;" json:"v_1"`                    //[ 3] v2                                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	V2    string `gorm:"column:v2;type:varchar;" json:"v_2"`                    //[ 4] v3                                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	V3    string `gorm:"column:v3;type:varchar;" json:"v_3"`                    //[ 5] v4                                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	V4    string `gorm:"column:v4;type:varchar;" json:"v_4"`                    //[ 6] v5                                             varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	V5    string `gorm:"column:v5;type:varchar;" json:"v_5"`
}

var casbin_ruleTableInfo = &TableInfo{
	Name: "casbin_rule",
	Columns: []*ColumnInfo{
		&ColumnInfo{
			Index:   0,
			Name:    "p_type",
			Comment: ``,
			Notes: `Warning table: casbin_rule does not have a primary key defined, setting col position 1 p_type as primary key
Warning table: casbin_rule primary key column p_type is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "PType",
			GoFieldType:        "string",
			JSONFieldName:      "p_type",
			ProtobufFieldName:  "p_type",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},
		&ColumnInfo{
			Index:              1,
			Name:               "v0",
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
			GoFieldName:        "V0",
			GoFieldType:        "string",
			JSONFieldName:      "v_0",
			ProtobufFieldName:  "v_0",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "v1",
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
			GoFieldName:        "V1",
			GoFieldType:        "string",
			JSONFieldName:      "v_1",
			ProtobufFieldName:  "v_1",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
			Name:               "v2",
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
			GoFieldName:        "V2",
			GoFieldType:        "string",
			JSONFieldName:      "v_2",
			ProtobufFieldName:  "v_2",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "v3",
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
			GoFieldName:        "V3",
			GoFieldType:        "string",
			JSONFieldName:      "v_3",
			ProtobufFieldName:  "v_3",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "v4",
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
			GoFieldName:        "V4",
			GoFieldType:        "string",
			JSONFieldName:      "v_4",
			ProtobufFieldName:  "v_4",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "v5",
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
			GoFieldName:        "V5",
			GoFieldType:        "string",
			JSONFieldName:      "v_5",
			ProtobufFieldName:  "v_5",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CasbinRule) TableName() string {
	return "casbin_rule"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CasbinRule) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CasbinRule) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CasbinRule) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CasbinRule) TableInfo() *TableInfo {
	return casbin_ruleTableInfo
}

// GetAllCasbinRule is a function to get a slice of record(s) from casbin_rule table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (c *CasbinRule) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*CasbinRule, totalRows int64, err error) {

	resultOrm := DB.Model(&CasbinRule{})
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

// GetCasbinRule is a function to get a single record from the casbin_rule table in the daigou database
// error - model.ErrNotFound, db Find error
func (c *CasbinRule) Get(DB *gorm.DB, argPType string) (record *CasbinRule, err error) {
	record = &CasbinRule{}
	if err = DB.First(record, argPType).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddCasbinRule is a function to add a single record to casbin_rule table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (c *CasbinRule) Add(DB *gorm.DB, record *CasbinRule) (result *CasbinRule, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCasbinRule is a function to update a single record from casbin_rule table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (c *CasbinRule) Update(DB *gorm.DB, argPType string, updated *CasbinRule) (result *CasbinRule, RowsAffected int64, err error) {

	result = &CasbinRule{}
	db := DB.First(result, argPType)
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

// DeleteCasbinRule is a function to delete a single record from casbin_rule table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (c *CasbinRule) Delete(DB *gorm.DB, argPType string) (rowsAffected int64, err error) {

	record := &CasbinRule{}
	db := DB.First(record, argPType)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
