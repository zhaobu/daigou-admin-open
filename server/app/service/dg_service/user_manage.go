package dg_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/model"
	"daigou-admin/model/dg_model"
)

func DeleteUserInfoByIds(ids request.IdsReq) (err error) {
	err = global.GetSerDB().Delete(&[]dg_model.User{}, "id in (?)", ids.Ids).Error
	return err
}

func FindUserInfoList(req *request.FindUserInfoListReq) (err error, resp *response.FindUserInfoListResp) {
	resp = &response.FindUserInfoListResp{
		TotalCount: 0,
		Info:       make([]*response.UserInfoList, 0, req.PageSize),
	}

	db := global.GetSerDB().Model(&dg_model.User{})

	//查询条件
	if req.PhoneNumber != "" {
		db = db.Where("user.phone_number=? ", req.PhoneNumber)
	}
	if req.NickName != "" {
		db = db.Where("user.nick_name=?", req.NickName)
	}
	//总分页数
	db.Count(&resp.TotalCount)

	//查询目标
	db = db.Select("user.*,shop_info.shop_id")

	//连接查询
	db = db.Joins("left join shop_info on user.user_id=shop_info.user_id")

	//排序顺序
	if req.Orders != "" {
		db = db.Order(req.Orders)
	}

	//处理分页
	if req.Page > 0 {
		offset := (req.Page - 1) * req.PageSize
		db = db.Offset(offset).Limit(req.PageSize)
	} else {
		db = db.Limit(req.PageSize)
	}

	//查询结果
	if err = db.Find(&resp.Info).Error; err != nil {
		err = model.ErrNotFound
		return err, resp
	}

	return err, resp
}

func OperateUser(req *request.OperateUserReq) (err error, resp *response.OperateUserResp, msg string) {
	resp = &response.OperateUserResp{}
	switch req.Operate {
	//开通店铺
	case request.OperateUser_OpenShop:

	}

	return
}
