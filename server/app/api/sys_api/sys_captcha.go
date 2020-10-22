package sys_api

import (
	"daigou-admin/config"
	"daigou-admin/global"
	"daigou-admin/global/response"
	"daigou-admin/utils/respwrite"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// @Tags base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.G_Config.Captcha.ImgHeight, global.G_Config.Captcha.ImgWidth, global.G_Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	var result string
	if global.G_Config.System.Mode == config.Debug {
		result = cp.Store.Get(id, false)
	}
	if err != nil {
		respwrite.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		respwrite.OkDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
			Result:    result,
		}, "验证码获取成功", c)
	}
}
