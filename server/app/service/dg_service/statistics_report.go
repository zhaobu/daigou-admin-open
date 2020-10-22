package dg_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/global/response"
	"fmt"
)

func NewUser(req *request.NewUserReq) (err error, resp []*response.NewUserResp) {

	if req.EndTime.Before(req.StartTime) {
		err = fmt.Errorf("结束时间不能早于开始时间")
		return
	}
	tb := global.GetSerDB().Table("user")
	resp = make([]*response.NewUserResp, 0, 10)
	if req.QueryType == request.Day {
		tb.Select("Date(created_at) as register_time,count(user_id) as num").Where("created_at between ? and ?", req.StartTime, req.EndTime).Group("register_time").Order("register_time").Scan(&resp)
	} else if req.QueryType == request.Week {
		tb.Select("concat(date_format(created_at,'%Y-'),week(created_at)) as register_time,count(user_id) as num").Where("created_at between ? and ?", req.StartTime, req.EndTime).Group("register_time").Order("register_time").Scan(&resp)
	} else if req.QueryType == request.Month {
		tb.Select("date_format(created_at,'%Y-%m') as register_time,count(user_id) as num").Where("created_at between ? and ?", req.StartTime, req.EndTime).Group("register_time").Order("register_time").Scan(&resp)
	}
	return
}

func ActiveUser(req *request.ActiveUserReq) (err error, resp []*response.ActiveUserResp) {

	if req.EndTime.Before(req.StartTime) {
		err = fmt.Errorf("结束时间不能早于开始时间")
		return
	}
	tb := global.GetSerDB().Table("user")
	resp = make([]*response.ActiveUserResp, 0, 10)
	if req.QueryType == request.Day {
		tb.Select("Date(updated_at) as register_time,count(user_id) as num").Where("updated_at between ? and ?", req.StartTime, req.EndTime).Group("register_time").Order("register_time").Scan(&resp)
	} else if req.QueryType == request.Week {
		tb.Select("concat(date_format(updated_at,'%Y-'),week(updated_at)) as register_time,count(user_id) as num").Where("updated_at between ? and ?", req.StartTime, req.EndTime).Group("register_time").Order("register_time").Scan(&resp)
	} else if req.QueryType == request.Month {
		tb.Select("date_format(updated_at,'%Y-%m') as register_time,count(user_id) as num").Where("updated_at between ? and ?", req.StartTime, req.EndTime).Group("register_time").Order("register_time").Scan(&resp)
	}
	return
}
