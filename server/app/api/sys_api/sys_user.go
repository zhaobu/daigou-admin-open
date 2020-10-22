package sys_api

import (
	"daigou-admin/app/service/sys_service"
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/middleware"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils"
	"daigou-admin/utils/respwrite"
	"daigou-admin/utils/zaplog"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body sys_model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"NickName":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		respwrite.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user := &sys_model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := sys_service.Register(*user)
	if err != nil {
		respwrite.FailWithDetailed(respwrite.ERROR, response.SysUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		respwrite.OkDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"CaptchaId": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		respwrite.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &sys_model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := sys_service.Login(U); err != nil {
			respwrite.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		respwrite.FailWithMessage("验证码错误", c)
	}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user sys_model.SysUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.G_Config.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "qmPlus",                       // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		respwrite.FailWithMessage("获取token失败", c)
		return
	}
	if !global.G_Config.System.UseMultipoint {
		respwrite.OkWithData(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
		return
	}
	var loginJwt sys_model.JwtBlacklist
	loginJwt.Jwt = token
	err, jwtStr := sys_service.GetRedisJWT(user.Username)
	if err == redis.Nil {
		if err := sys_service.SetRedisJWT(loginJwt, user.Username); err != nil {
			respwrite.FailWithMessage("设置登录状态失败", c)
			return
		}
		respwrite.OkWithData(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	} else if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		var blackJWT sys_model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := sys_service.JsonInBlacklist(blackJWT); err != nil {
			respwrite.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := sys_service.SetRedisJWT(loginJwt, user.Username); err != nil {
			respwrite.FailWithMessage("设置登录状态失败", c)
			return
		}
		data := &response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}
		zaplog.Debugf("tokenNext:%s", utils.Dump(data))
		respwrite.OkWithData(data, c)
	}
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(c *gin.Context) {
	var params request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&params)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"NewPassword": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		respwrite.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	U := &sys_model.SysUser{Username: params.Username, Password: params.Password}
	if err, _ := sys_service.ChangePassword(U, params.NewPassword); err != nil {
		respwrite.FailWithMessage("修改失败，请检查用户名密码", c)
	} else {
		respwrite.OkWithMessage("修改成功", c)
	}
}

type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// @Tags SysUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*request.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	// 便于找到用户 以后从jwt中取
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		// 文件上传后拿到文件路径
		err, filePath, _ := utils.Upload(header)
		if err != nil {
			respwrite.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			err, user := sys_service.UploadHeaderImg(uuid, filePath)
			if err != nil {
				respwrite.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				respwrite.OkWithData(response.SysUserResponse{User: *user}, c)
			}
		}
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		respwrite.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := sys_service.GetUserInfoList(pageInfo)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	UserVerify := utils.Rules{
		"UUID":        {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(sua, UserVerify)
	if UserVerifyErr != nil {
		respwrite.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	err := sys_service.SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("修改失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		respwrite.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := sys_service.DeleteUser(reqId.Id)
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		respwrite.OkWithMessage("删除成功", c)
	}
}
