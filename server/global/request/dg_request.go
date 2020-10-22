package request

import (
	"time"
)

/***************************server_config****************************/
type ChangeConnectReq struct {
	Index int `form:"index" v:"index@required#index不能为空"`
}

// type ChangeRedisReq struct {
// 	Index int `form:"index" v:"index@required#index不能为空"`
// }

// type ChangeMysqlReq struct {
// 	Index int `form:"index" v:"index@required#index不能为空"`
// }

/***************************statics_report****************************/
type QueryType uint32

const (
	Day QueryType = iota
	Week
	Month
)

// 每日新增用户
type NewUserReq struct {
	QueryType QueryType `form:"query_type" v:"query_type@required#query_type不能为空"` //查询类型
	StartTime time.Time `form:"start_time" v:"start_time@required#start_time不能为空"`
	EndTime   time.Time `form:"end_time" v:"end_time@required#end_time不能为空"`
}

// 每日新增用户
type ActiveUserReq struct {
	QueryType QueryType `form:"query_type" v:"query_type@required#query_type不能为空"` //查询类型
	StartTime time.Time `form:"start_time" v:"start_time@required#start_time不能为空"`
	EndTime   time.Time `form:"end_time" v:"end_time@required#end_time不能为空"`
}

/***************************shop_manage****************************/
type FindShopInfoListReq struct {
	PageInfo
	ShopName    string `json:"shopName" `    //店铺名称
	PhoneNumber string `json:"phoneNumber" ` //手机号码
	Orders      string `json:"orders" `      //排序顺序
}

type OperateShop int32

const (
	OperateShop_None         OperateShop = iota - 1
	OperateShop_FreezeShop               //冻结店铺
	OperateShop_UnFreezeShop             //解冻店铺
	OperateShop_CloseShop                //关闭店铺
	OperateShop_OpenShop                 //开启店铺
	OperateShop_UpgradeShop              //升级店铺
)

type OperateShopReq struct {
	Operate OperateShop `form:"operate" json:"operate" v:"operate @required|between:0,4 #不能为空|操作类型未定义"`               //操作类型
	ShopID  uint64      `form:"shop_id" json:"shop_id" v:"shop_id @required|min-length:7 #shop_id不能为空|shop_id最短长度为7"` // 店铺ID
}

/***************************user_manage****************************/
type FindUserInfoListReq struct {
	PageInfo
	NickName    string `json:"nickName" `    //微信昵称
	PhoneNumber string `json:"phoneNumber" ` //手机号码
	Orders      string `json:"orders" `      //排序顺序
}

type OperateUser int32

const (
	OperateUser_None     OperateUser = iota - 1
	OperateUser_OpenShop             //开通店铺
)

type OperateUserReq struct {
	Operate OperateUser `form:"operate" json:"operate" v:"operate @required|between:0,4 #不能为空|操作类型未定义"`               //操作类型
	UserID  uint64      `form:"user_id" json:"user_id" v:"user_id @required|min-length:7 #user_id不能为空|user_id最短长度为7"` // 店铺ID
	Args    string      `form:"args" json:"args" v:"args@required #args不能为空"`                                         // 参数
}

/***************************proxy_api****************************/
type ProxyPostReq struct {
	URL  string      `json:"url" `
	Data interface{} `json:"data" `
}
