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


CREATE TABLE `vip_recharge_records` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `recharge_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '充值流水编号',
  `trand_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '商户流水编号',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `member_price` decimal(10,2) unsigned NOT NULL COMMENT '充值金额',
  `fee_type` int(10) unsigned NOT NULL COMMENT '充值类型（1：一个月 2：一个季度 3：一年 4：自定义月数）',
  `status` int(11) NOT NULL COMMENT '订单状态0发起1成功-1异常',
  `remark` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '订单描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`,`trand_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='会员充值流水表'


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// VipRechargeRecords struct is a row record of the vip_recharge_records table in the daigou database
type VipRechargeRecords struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID          uint32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 主键//[ 1] recharge_id                                    varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	RechargeID  string      `gorm:"column:recharge_id;type:varchar;" json:"recharge_id"`       // 充值流水编号//[ 2] trand_id                                       varchar              null: false  primary: true   isArray: false  auto: false  col: varchar         len: 0       default: []
	TrandID     string      `gorm:"primary_key;column:trand_id;type:varchar;" json:"trand_id"` // 商户流水编号//[ 3] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID      uint32      `gorm:"column:user_id;type:uint;" json:"user_id"`                  // 用户id//[ 4] member_price                                   udecimal             null: false  primary: false  isArray: false  auto: false  col: udecimal        len: -1      default: []
	MemberPrice float64     `gorm:"column:member_price;type:udecimal;" json:"member_price"`    // 充值金额//[ 5] fee_type                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	FeeType     uint32      `gorm:"column:fee_type;type:uint;" json:"fee_type"`                // 充值类型（1：一个月 2：一个季度 3：一年 4：自定义月数）//[ 6] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Status      int32       `gorm:"column:status;type:int;" json:"status"`                     // 订单状态0发起1成功-1异常//[ 7] remark                                         varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Remark      string      `gorm:"column:remark;type:varchar;" json:"remark"`                 // 订单描述//[ 8] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt   *model.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`        //[ 9] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt   *model.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var vip_recharge_recordsTableInfo = &TableInfo{
	Name: "vip_recharge_records",
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
			Name:               "recharge_id",
			Comment:            `充值流水编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "RechargeID",
			GoFieldType:        "string",
			JSONFieldName:      "recharge_id",
			ProtobufFieldName:  "recharge_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
		&ColumnInfo{
			Index:              2,
			Name:               "trand_id",
			Comment:            `商户流水编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "TrandID",
			GoFieldType:        "string",
			JSONFieldName:      "trand_id",
			ProtobufFieldName:  "trand_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},
		&ColumnInfo{
			Index:              4,
			Name:               "member_price",
			Comment:            `充值金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "udecimal",
			DatabaseTypePretty: "udecimal",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "udecimal",
			ColumnLength:       -1,
			GoFieldName:        "MemberPrice",
			GoFieldType:        "float64",
			JSONFieldName:      "member_price",
			ProtobufFieldName:  "member_price",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
		&ColumnInfo{
			Index:              5,
			Name:               "fee_type",
			Comment:            `充值类型（1：一个月 2：一个季度 3：一年 4：自定义月数）`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "FeeType",
			GoFieldType:        "uint32",
			JSONFieldName:      "fee_type",
			ProtobufFieldName:  "fee_type",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},
		&ColumnInfo{
			Index:              6,
			Name:               "status",
			Comment:            `订单状态0发起1成功-1异常`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},
		&ColumnInfo{
			Index:              7,
			Name:               "remark",
			Comment:            `订单描述`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       0,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
		&ColumnInfo{
			Index:              8,
			Name:               "created_at",
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
			GoFieldName:        "CreatedAt",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},
		&ColumnInfo{
			Index:              9,
			Name:               "updated_at",
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
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "*model.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (v *VipRechargeRecords) TableName() string {
	return "vip_recharge_records"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (v *VipRechargeRecords) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (v *VipRechargeRecords) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (v *VipRechargeRecords) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (v *VipRechargeRecords) TableInfo() *TableInfo {
	return vip_recharge_recordsTableInfo
}

// GetAllVipRechargeRecords is a function to get a slice of record(s) from vip_recharge_records table in the daigou database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - model.ErrNotFound, db Find error
func (v *VipRechargeRecords) GetAll(DB *gorm.DB, page, pagesize int, order string) (results []*VipRechargeRecords, totalRows int64, err error) {

	resultOrm := DB.Model(&VipRechargeRecords{})
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

// GetVipRechargeRecords is a function to get a single record from the vip_recharge_records table in the daigou database
// error - model.ErrNotFound, db Find error
func (v *VipRechargeRecords) Get(DB *gorm.DB, argId uint32, argTrandID string) (record *VipRechargeRecords, err error) {
	record = &VipRechargeRecords{}
	if err = DB.First(record, argId, argTrandID).Error; err != nil {
		err = model.ErrNotFound
		return record, err
	}

	return record, nil
}

// AddVipRechargeRecords is a function to add a single record to vip_recharge_records table in the daigou database
// error - model.ErrInsertFailed, db save call failed
func (v *VipRechargeRecords) Add(DB *gorm.DB, record *VipRechargeRecords) (result *VipRechargeRecords, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, model.ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateVipRechargeRecords is a function to update a single record from vip_recharge_records table in the daigou database
// error - model.ErrNotFound, db record for id not found
// error - model.ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (v *VipRechargeRecords) Update(DB *gorm.DB, argId uint32, argTrandID string, updated *VipRechargeRecords) (result *VipRechargeRecords, RowsAffected int64, err error) {

	result = &VipRechargeRecords{}
	db := DB.First(result, argId, argTrandID)
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

// DeleteVipRechargeRecords is a function to delete a single record from vip_recharge_records table in the daigou database
// error - model.ErrNotFound, db Find error
// error - model.ErrDeleteFailed, db Delete failed error
func (v *VipRechargeRecords) Delete(DB *gorm.DB, argId uint32, argTrandID string) (rowsAffected int64, err error) {

	record := &VipRechargeRecords{}
	db := DB.First(record, argId, argTrandID)
	if db.Error != nil {
		return -1, model.ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, model.ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
