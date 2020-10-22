package request

import "daigou-admin/config"

type GetDBReq struct {
	ConnectType config.DBType `form:"connectType" ` //数据库类型
}

type GetTablesReq struct {
	ConnectType config.DBType `form:"connectType" ` //数据库类型
	DBName      string        `form:"dbName" `      //数据库名称
}

type GetColumeReq struct {
	ConnectType config.DBType `form:"connectType" ` //数据库类型
	DBName      string        `form:"dbName" `      //数据库名称
	TableName   string        `form:"tableName" `   //表名称
}
