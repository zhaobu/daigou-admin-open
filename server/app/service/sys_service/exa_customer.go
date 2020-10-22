package sys_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/model/sys_model"
)

// @title    CreateExaCustomer
// @description   create a customer, 创建用户
// @param     e               sys_model.ExaCustomer
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateExaCustomer(e sys_model.ExaCustomer) (err error) {
	err = global.GetSysDB().Create(&e).Error
	return err
}

// @title    DeleteFileChunk
// @description   delete a customer, 删除用户
// @auth                     （2020/04/05  20:22）
// @param     e               sys_model.ExaCustomer
// @return                    error

func DeleteExaCustomer(e sys_model.ExaCustomer) (err error) {
	err = global.GetSysDB().Delete(e).Error
	return err
}

// @title    UpdateExaCustomer
// @description   update a customer, 更新用户
// @param     e               *sys_model.ExaCustomer
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateExaCustomer(e *sys_model.ExaCustomer) (err error) {
	err = global.GetSysDB().Save(e).Error
	return err
}

// @title    GetExaCustomer
// @description   get the info of a costumer , 获取用户信息
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    customer        ExaCustomer

func GetExaCustomer(id uint) (err error, customer sys_model.ExaCustomer) {
	err = global.GetSysDB().Where("id = ?", id).First(&customer).Error
	return
}

// @title    GetCustomerInfoList
// @description   get customer list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     sysUserAuthorityID              string
// @param     info            PageInfo
// @return                    error

func GetCustomerInfoList(sysUserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GetSysDB().Model(&sys_model.ExaCustomer{})
	var a sys_model.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	err, auth := GetAuthorityInfo(a)
	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []sys_model.ExaCustomer
	err = db.Where("sys_user_authority_id in (?)", dataId).Count(&total).Error
	if err != nil {
		return err, CustomerList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in (?)", dataId).Find(&CustomerList).Error
	}
	return err, CustomerList, total
}
