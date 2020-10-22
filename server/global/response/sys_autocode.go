package response

import "daigou-admin/config"

type DBResp struct {
	Database string `json:"database" gorm:"column:database"`
}

type TableResp struct {
	TableName string `json:"tableName"`
}

type ColumeResp struct {
	ColumeName    string `json:"columeName" gorm:"column:colume_name"`
	DataType      string `json:"dataType" gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumeComment string `json:"columeComment" gorm:"column:colume_comment"`
}

type Connect struct {
	Name        string        `json:"name" `        //配置名称
	DBName      string        `json:"dbName" `      //数据库名称
	ConnectType config.DBType `json:"connectType" ` //数据库类型
}

type DBConnectResp struct {
	Connects map[int]*Connect `json:"connects" `
}
