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


CREATE TABLE `test_db` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `type_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `name` varchar(255) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `avatar_url` varchar(255) DEFAULT '',
  `passwd` varchar(200) NOT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `gender` enum('one','two') CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `price` decimal(10,2) unsigned DEFAULT NULL,
  `detial` json DEFAULT NULL,
  PRIMARY KEY (`id`,`type_id`) USING BTREE,
  UNIQUE KEY `uix_test_db_email` (`email`) USING BTREE,
  KEY `idx_test_db_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 4] column is set for unsigned
[14] column is set for unsigned



*/

// TestDb struct is a row record of the test_db table in the daigou database
type TestDb struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID        uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`         // 主键//[ 1] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`                // 创建时间//[ 2] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`                // 更新时间//[ 3] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt *model.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`                // 删除时间//[ 4] type_id                                        ubigint              null: false  primary: true   isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TypeID    uint64      `gorm:"primary_key;column:type_id;type:ubigint;default:0;" json:"type_id"` //[ 5] name                                           varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Name      string      `gorm:"column:name;type:varchar;" json:"name"`                             //[ 6] age                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Age       int32       `gorm:"column:age;type:int;" json:"age"`                                   //[ 7] avatar_url                                     varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	AvatarURL string      `gorm:"column:avatar_url;type:varchar;" json:"avatar_url"`                 //[ 8] passwd                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Passwd    string      `gorm:"column:passwd;type:varchar;" json:"passwd"`                         //[ 9] created                                        datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Created   *model.Time `gorm:"column:created;type:datetime;" json:"created"`                      //[10] updated                                        datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Updated   *model.Time `gorm:"column:updated;type:datetime;" json:"updated"`                      //[11] gender                                         char                 null: true   primary: false  isArray: false  auto: false  col: char            len: 0       default: []
	Gender    string      `gorm:"column:gender;type:char;" json:"gender"`                            //[12] birthday                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Birthday  *model.Time `gorm:"column:birthday;type:datetime;" json:"birthday"`                    //[13] email                                          varchar              null: true   primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Email     string      `gorm:"column:email;type:varchar;" json:"email"`                           //[14] price                                          udecimal             null: true   primary: false  isArray: false  auto: false  col: udecimal        len: -1      default: []
	Price     float64     `gorm:"column:price;type:udecimal;" json:"price"`                          //[15] detial                                         json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Detial    string      `gorm:"column:detial;type:json;" json:"detial"`
}

var test_dbTableInfo = &TableInfo{
	Name: "test_db",
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
			Name:               "created_at",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
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
			Nullable:           false,
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
			Name:               "type_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TypeID",
			GoFieldType:        "uint64",
			JSONFieldName:      "type_id",
			ProtobufFieldName:  "type_id",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "age",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Age",
			GoFieldType:        "int32",
			JSONFieldName:      "age",
			ProtobufFieldName:  "age",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "avatar_url",
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
			GoFieldName:        "AvatarURL",
			GoFieldType:        "string",
			JSONFieldName:      "avatar_url",
			ProtobufFieldName:  "avatar_url",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "passwd",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Passwd",
			GoFieldType:        "string",
			JSONFieldName:      "passwd",
			ProtobufFieldName:  "passwd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "created",
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
			GoFieldName:        "Created",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "created",
			ProtobufFieldName:  "created",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},
		&ColumnInfo{
			Index:              10,
			Name:               "updated",
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
			GoFieldName:        "Updated",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "updated",
			ProtobufFieldName:  "updated",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},
		&ColumnInfo{
			Index:              11,
			Name:               "gender",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       0,
			GoFieldName:        "Gender",
			GoFieldType:        "string",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
		&ColumnInfo{
			Index:              12,
			Name:               "birthday",
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
			GoFieldName:        "Birthday",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "birthday",
			ProtobufFieldName:  "birthday",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        13,
		},
		&ColumnInfo{
			Index:              13,
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},
		&ColumnInfo{
			Index:              14,
			Name:               "price",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "udecimal",
			DatabaseTypePretty: "udecimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "udecimal",
			ColumnLength:       -1,
			GoFieldName:        "Price",
			GoFieldType:        "float64",
			JSONFieldName:      "price",
			ProtobufFieldName:  "price",
			ProtobufType:       "float",
			ProtobufPos:        15,
		},
		&ColumnInfo{
			Index:              15,
			Name:               "detial",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "Detial",
			GoFieldType:        "string",
			JSONFieldName:      "detial",
			ProtobufFieldName:  "detial",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TestDb) TableName() string {
	return "test_db"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TestDb) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TestDb) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TestDb) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TestDb) TableInfo() *TableInfo {
	return test_dbTableInfo
}

// GetAllTestDb is a function to get a slice of record(s) from test_db table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (t *TestDb) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*TestDb, totalRows int64, err error) {

	resultOrm := DB.Model(&TestDb{})
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

// GetTestDb is a function to get a single record from the test_db table in the daigou database
// error - model.ErrNotFound, db Find error
func (t *TestDb) Get(DB *gorm.DB, argId uint32, argTypeID uint64) (record *TestDb, err error) {
	record = &TestDb{}
	if err = DB.First(record, argId, argTypeID).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddTestDb is a function to add a single record to test_db table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (t *TestDb) Add(DB *gorm.DB, record *TestDb) (result *TestDb, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateTestDb is a function to update a single record from test_db table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (t *TestDb) Update(DB *gorm.DB, argId uint32, argTypeID uint64, updated *TestDb) (result *TestDb, RowsAffected int64, err error) {

	result = &TestDb{}
	db := DB.First(result, argId, argTypeID)
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

// DeleteTestDb is a function to delete a single record from test_db table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (t *TestDb) Delete(DB *gorm.DB, argId uint32, argTypeID uint64) (rowsAffected int64, err error) {

	record := &TestDb{}
	db := DB.First(record, argId, argTypeID)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
