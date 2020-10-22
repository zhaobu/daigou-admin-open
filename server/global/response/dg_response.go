package response

import "daigou-admin/model/dg_model"

//封装有多种返回结果的情况,不会传给客户端
type BaseResp struct {
	Code int    `json:"-"`
	Msg  string `json:"-"`
}

/***************************server_config****************************/
type GetConnectResp struct {
	CurrentIndex    int                  `json:"current_index" `
	RedisConfigList map[int]*RedisConfig `json:"redis_config_list" `
	MysqlConfigList map[int]*MysqlConfig `json:"mysql_config_list" `
}

type RedisConfig struct {
	Addr     string `json:"addr" `
	Password string `json:"password" `
	DB       int    `json:"db" `
	Name     string `json:"name" `
	Index    int    `json:"index" `
}

// type GetRedisListResp struct {
// 	CurrentIndex    int                  `json:"current_index" `
// 	RedisConfigList map[int]*RedisConfig `json:"redis_config_list" `
// }

type MysqlConfig struct {
	Name     string            `json:"name" `      //配置名称
	UserName string            `json:"user_name" ` //用户名
	Password string            `json:"password" `  //密码
	Addr     string            `json:"addr" `      //地址
	DBName   string            `json:"db_name" `   //数据库名称
	Params   map[string]string `json:"params" `    //数据库连接参数
	Index    int               `json:"index" `
}

// type GetMysqlListResp struct {
// 	CurrentIndex    int                  `json:"current_index" `
// 	MysqlConfigList map[int]*MysqlConfig `json:"mysql_config_list" `
// }

/***************************statics_report****************************/
type NewUserResp struct {
	RegisterTime string `json:"register_time" `
	Num          uint32 `json:"num" `
}

type ActiveUserResp struct {
	RegisterTime string `json:"register_time" `
	Num          uint32 `json:"num" `
}

/***************************shop_manage****************************/
type ShopInfoList struct {
	*dg_model.ShopInfo
	PhoneNumber string `gorm:"column:phone_number;type:varchar;" json:"phone_number"` // 手机号码
}
type FindShopInfoListResp struct {
	Info       []*ShopInfoList
	TotalCount int
}

type OperateShopResp struct {
}

/***************************user_manage****************************/
type UserInfoList struct {
	*dg_model.User
	ShopID uint64 `gorm:"column:shop_id;type:ubigint;" json:"shop_id"` // 店铺ID
}
type FindUserInfoListResp struct {
	Info       []*UserInfoList
	TotalCount int
}

type OperateUserResp struct {
	Data interface{} `json:"data"`
}

/***************************proxy_api****************************/
type RespCode int32

const (
	ERROR         RespCode = iota - 1 //请求错误
	SUCCESS                           //请求成功
	FAIL                              //请求失败,一般指业务方面
	UNAUTHORIZED                      //鉴权失败
	EXPIRED                           //token失效(只在访问登录接口时使用)
	UNAUTHOPERATE                     //没有操作权限
)

type ProxyPostResp struct {
	Code RespCode    `json:"code"` // 错误码(-1:请求错误 0:请求成功 1:请求失败,一般指业务方面 2:鉴权失败,3:token失效 4:没有操作权限 )
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}
