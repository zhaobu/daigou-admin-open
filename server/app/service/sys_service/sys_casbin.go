package sys_service

import (
	"daigou-admin/global"
	"daigou-admin/global/request"
	"daigou-admin/model/sys_model"
	"daigou-admin/utils/zaplog"
	"errors"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

// @title    UpdateCasbin
// @description   update casbin authority, 更新casbin权限
// @auth                     （2020/04/05  20:22）
// @param     authorityId      string
// @param     casbinInfos      []CasbinInfo
// @return                     error

func UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	ClearCasbin(0, authorityId)
	for _, v := range casbinInfos {
		cm := sys_model.CasbinModel{
			ID:          0,
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		addflag := AddCasbin(cm)
		if addflag == false {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

// @title    AddCasbin
// @description   add casbin authority, 添加权限
// @auth                     （2020/04/05  20:22）
// @param     cm              sys_model.CasbinModel
// @return                    bool

func AddCasbin(cm sys_model.CasbinModel) bool {
	e := Casbin()
	res, err := e.AddPolicy(cm.AuthorityId, cm.Path, cm.Method)
	if err != nil {
		zaplog.Errorf("AddCasbin err:%s", err)
	}
	return res
}

// @title    UpdateCasbinApi
// @description   update casbin apis, API更新随动
// @auth                     （2020/04/05  20:22）
// @param     oldPath          string
// @param     newPath          string
// @param     oldMethod        string
// @param     newMethod        string
// @return                     error

func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	var cs []sys_model.CasbinModel
	err := global.GetSysDB().Table("casbin_rule").Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Find(&cs).Updates(map[string]string{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// @title    GetPolicyPathByAuthorityId
// @description   get policy path by authorityId, 获取权限列表
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return                    []string

func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// @title    ClearCasbin
// @description   清除匹配的权限
// @auth                     （2020/04/05  20:22）
// @param     v               int
// @param     p               string
// @return                    bool

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	res, err := e.RemoveFilteredPolicy(v, p...)
	if err != nil {
		zaplog.Errorf("ClearCasbin err:%s", err)
	}
	return res
}

// @title    Casbin
// @description   store to DB, 持久化到数据库  引入自定义规则
// @auth                     （2020/04/05  20:22）

func Casbin() *casbin.CachedEnforcer {
	e := global.G_Casbin.GetEnforcer()
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	return e
}

// @title    ParamsMatch
// @description   customized rule, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     fullNameKey1    string
// @param     key2            string
// @return                    bool

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @title    ParamsMatchFunc
// @description   customized function, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     args            ...interface{}
// @return                    interface{}
// @return                    error

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
