package dg_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"daigou-admin/utils/zaplog"
	"fmt"

	"github.com/gogf/gf/net/ghttp"
	jsoniter "github.com/json-iterator/go"
)

//代理访问

func ProxyPost(req *request.ProxyPostReq) (err error, resp *response.ProxyPostResp, msg string) {
	var (
		baseURL = global.G_Config.Service[0].BaseURL
		dgResp  *ghttp.ClientResponse
	)
	resp = &response.ProxyPostResp{}

	defer func() {
		if err != nil {
			zaplog.Errorf("err=%s", err)
		}
		if dgResp != nil {
			dgResp.Close()
		}
	}()
	switch req.URL {
	case "admin/getShopCode":
		dgResp, err = ghttp.Post(baseURL+req.URL, req.Data)
	case "admin/useShopCode":
		dgResp, err = ghttp.Post(baseURL+req.URL, req.Data)
	default:
		err = fmt.Errorf("未定义的请求")
	}
	if err != nil {
		return err, resp, msg
	}
	//解析请求结果
	if err = jsoniter.Unmarshal(dgResp.ReadAll(), resp); err != nil {
		return
	}
	//解析结果中的data
	return
}
