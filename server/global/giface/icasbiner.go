package giface

import "github.com/casbin/casbin/v2"

type Casbiner interface {
	Setup()
	GetEnforcer() *casbin.CachedEnforcer
}
