package authcasbin

import (
	"daigou-admin/global"
	"daigou-admin/global/giface"
	"daigou-admin/utils/zaplog"

	"github.com/casbin/casbin/v2/model"

	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && ParamsMatch(r.obj,p.obj) && r.act == p.act
`

type Casbin struct {
	giface.Casbiner
	CachedEnforcer *casbin.CachedEnforcer
}

func NewCasbin() *Casbin {
	return &Casbin{}
}

func (c *Casbin) Setup() {
	Apter, err := gormAdapter.NewAdapterByDB(global.GetSysDB().DB)
	if err != nil {
		zaplog.Panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		zaplog.Panic(err)
	}
	c.CachedEnforcer, err = casbin.NewCachedEnforcer(m, Apter)
	if err != nil {
		zaplog.Panic(err)
	}
	err = c.CachedEnforcer.LoadPolicy()
	if err != nil {
		zaplog.Panic(err)
	}
}

func (c *Casbin) GetEnforcer() *casbin.CachedEnforcer {
	return c.CachedEnforcer
}
