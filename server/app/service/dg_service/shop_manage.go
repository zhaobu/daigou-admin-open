package dg_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/model"
	"daigou-admin/model/dg_model"
	"fmt"
)

func DeleteShopInfoByIds(ids request.IdsReq) (err error) {
	err = global.GetSerDB().Delete(&[]dg_model.ShopInfo{}, "id in (?)", ids.Ids).Error
	return err
}

func FindShopInfoList(req *request.FindShopInfoListReq) (err error, resp *response.FindShopInfoListResp) {
	resp = &response.FindShopInfoListResp{
		TotalCount: 0,
		Info:       make([]*response.ShopInfoList, 0, req.PageSize),
	}

	db := global.GetSerDB().Model(&dg_model.ShopInfo{})

	//查询条件
	if req.PhoneNumber != "" {
		db = db.Where("user.phone_number=? ", req.PhoneNumber)
	}
	if req.ShopName != "" {
		db = db.Where("shop_info.shop_name=?", req.ShopName)
	}

	//总分页数
	db.Count(&resp.TotalCount)

	//查询目标
	db = db.Select("shop_info.*,user.phone_number")

	//连接查询
	db = db.Joins("left join user on user.user_id=shop_info.user_id")

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

func OperateShop(req *request.OperateShopReq) (err error, resp *response.OperateShopResp, msg string) {
	resp = &response.OperateShopResp{}
	db := global.GetSerDB().Model(&dg_model.ShopInfo{})
	switch req.Operate {
	case request.OperateShop_FreezeShop:
		err = db.Where("shop_id=?", req.ShopID).Updates(map[string]interface{}{"is_enable": 0}).Error
		msg = fmt.Sprintf("店铺%d冻结成功", req.ShopID)
	case request.OperateShop_UnFreezeShop:
		err = db.Where("shop_id=?", req.ShopID).Updates(map[string]interface{}{"is_enable": 1}).Error
		msg = fmt.Sprintf("店铺%d解冻成功", req.ShopID)
	default:
		err = fmt.Errorf("操作未定义")
	}
	if err != nil {
		msg = ""
		return
	}
	return
}
