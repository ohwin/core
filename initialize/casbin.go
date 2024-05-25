package initialize

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/ohwin/core/global"
	"github.com/ohwin/core/log"
)

func Casbin() error {
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		log.Warn("casbin model create error: %v", err)
		panic(err)
	}

	config := global.Config.Mysql
	a, err := adapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", config.User, config.Password, config.Host)) // Your driver and data source.
	if err != nil {
		log.Warn("casbin new adapter error: %v", err)
		panic(err)
	}

	global.Enforce, err = casbin.NewSyncedCachedEnforcer(m, a)
	if err != nil {
		log.Warn("casbin enforce error: %v", err)
		panic(err)
	}
	return nil
}
