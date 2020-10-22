package dg_model

import (
	"errors"
	"fmt"
	"reflect"
)

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["casbin_api"] = casbin_apiTableInfo
	tables["casbin_role_api"] = casbin_role_apiTableInfo
	tables["casbin_rule"] = casbin_ruleTableInfo
	tables["casbin_user_role"] = casbin_user_roleTableInfo
	tables["express_template_config"] = express_template_configTableInfo
	tables["goods"] = goodsTableInfo
	tables["goods_access_records"] = goods_access_recordsTableInfo
	tables["goods_category"] = goods_categoryTableInfo
	tables["goods_message"] = goods_messageTableInfo
	tables["orders"] = ordersTableInfo
	tables["orders_bill_flow"] = orders_bill_flowTableInfo
	tables["orders_goods"] = orders_goodsTableInfo
	tables["orders_logistics"] = orders_logisticsTableInfo
	tables["orders_printing"] = orders_printingTableInfo
	tables["shop_access_records"] = shop_access_recordsTableInfo
	tables["shop_category"] = shop_categoryTableInfo
	tables["shop_code_records"] = shop_code_recordsTableInfo
	tables["shop_express"] = shop_expressTableInfo
	tables["shop_fans"] = shop_fansTableInfo
	tables["shop_info"] = shop_infoTableInfo
	tables["shop_profit"] = shop_profitTableInfo
	tables["shop_teamuser"] = shop_teamuserTableInfo
	tables["shop_vip"] = shop_vipTableInfo
	tables["shop_vip_explain"] = shop_vip_explainTableInfo
	tables["shop_vip_price"] = shop_vip_priceTableInfo
	tables["shop_wallet"] = shop_walletTableInfo
	tables["system_ann"] = system_annTableInfo
	tables["system_config"] = system_configTableInfo
	tables["system_express_company"] = system_express_companyTableInfo
	tables["test_db"] = test_dbTableInfo
	tables["third_party"] = third_partyTableInfo
	tables["third_party_interface"] = third_party_interfaceTableInfo
	tables["user"] = userTableInfo
	tables["user_address"] = user_addressTableInfo
	tables["user_login_records"] = user_login_recordsTableInfo
	tables["user_set_up"] = user_set_upTableInfo
	tables["vip_recharge_records"] = vip_recharge_recordsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}

// Copy a src struct into a destination struct
func Copy(dst interface{}, src interface{}) error {
	dstV := reflect.Indirect(reflect.ValueOf(dst))
	srcV := reflect.Indirect(reflect.ValueOf(src))

	if !dstV.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	if srcV.Type() != dstV.Type() {
		return errors.New("different types can be copied")
	}

	for i := 0; i < dstV.NumField(); i++ {
		f := srcV.Field(i)
		if !isZeroOfUnderlyingType(f.Interface()) {
			dstV.Field(i).Set(f)
		}
	}

	return nil
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
