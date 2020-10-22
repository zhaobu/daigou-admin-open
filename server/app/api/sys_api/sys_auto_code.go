package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global/request"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"daigou-admin/utils/zaplog"
	"fmt"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sys_model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(c *gin.Context) {
	var a sys_model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	AutoCodeVerify := utils.Rules{
		"Abbreviation": {utils.NotEmpty()},
		"StructName":   {utils.NotEmpty()},
		"PackageName":  {utils.NotEmpty()},
		"Fields":       {utils.NotEmpty()},
	}
	WKVerifyErr := utils.Verify(a, AutoCodeVerify)
	if WKVerifyErr != nil {
		respwrite.FailWithMessage(WKVerifyErr.Error(), c)
		return
	}
	if a.AutoCreateApiToSql {
		apiList := [6]sys_model.SysApi{
			{
				Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
				Description: "新增" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
				Description: "删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
				Description: "批量删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
				Description: "更新" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "PUT",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
				Description: "根据ID获取" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
				Description: "获取" + a.Description + "列表",
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
		}
		for _, v := range apiList {
			errC := sys_service.CreateApi(v)
			if errC != nil {
				c.Writer.Header().Add("success", "false")
				c.Writer.Header().Add("msg", url.QueryEscape(fmt.Sprintf("自动化创建失败，%v，请自行清空垃圾数据", errC)))
				return
			}
		}
	}
	err := sys_service.CreateTemp(a)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
		os.Remove("./ginvueadmin.zip")
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		os.Remove("./ginvueadmin.zip")
	}
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getTables [get]

func GetTables(c *gin.Context) {
	var (
		req request.GetTablesReq
		err error
	)

	//获取参数
	if err = c.ShouldBindQuery(&req); err != nil {
		zaplog.Errorf("api GetTables err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("GetTables req:%s", utils.Dump(req))

	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api GetTables err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}
	err, tables := sys_service.GetTables(&req)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询table失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{
			"tables": tables,
		}, c)
	}
}

// @Tags SysApi
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDB [get]
func GetDB(c *gin.Context) {
	var (
		req request.GetDBReq
		err error
	)
	//获取参数
	if err = c.ShouldBindQuery(&req); err != nil {
		zaplog.Errorf("api GetDB err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("GetDB req:%s", utils.Dump(req))
	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api GetDB err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	err, dbs := sys_service.GetDB(&req)
	// zaplog.Debugf("GetDB resp:%s", utils.Dump(dbs))

	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询DB失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{
			"dbs": dbs,
		}, c)
	}
}

// @Tags SysApi
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getColume [get]
func GetColume(c *gin.Context) {
	var (
		req request.GetColumeReq
		err error
	)

	//获取参数
	if err = c.ShouldBindQuery(&req); err != nil {
		zaplog.Errorf("api GetColume err:%w", err)
		respwrite.FailWithMessage(err.Error(), c)
		return
	}
	zaplog.Debugf("GetColume req:%s", utils.Dump(req))
	//检查参数
	if err = gvalid.CheckStruct(&req, nil); err != (*gvalid.Error)(nil) {
		if v, ok := err.(*gvalid.Error); ok {
			zaplog.Errorf("api GetColume err:%s", v.FirstString())
			respwrite.FailWithMessage(v.FirstString(), c)
		}
		return
	}

	err, columes := sys_service.GetColume(&req)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询Colume失败，%v", err), c)
	} else {
		respwrite.OkWithData(gin.H{
			"columes": columes,
		}, c)
	}
}

// @Tags SysApi
// @Summary 获取当前业务数据库和后台数据库连接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDBConnect [get]
func GetDBConnect(c *gin.Context) {
	err, resp := sys_service.GetDBConnect()
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("查询DBConnect失败，%v", err), c)
	} else {
		respwrite.OkWithData(resp, c)
	}
}
