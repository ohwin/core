package core

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/global"
	"github.com/ohwin/core/initialize"
	"github.com/ohwin/core/log"
	"github.com/ohwin/core/types"
)

func Init() {
	initialize.Viper()
	initialize.Log()
	initialize.Redis()
	//initialize.MQ()
	initialize.DB()
	casbin := global.Config.Casbin
	if casbin {
		initialize.Casbin()
	}

}

func RunWindowsServer(routers []types.RouterFunc, inits []func()) {
	// 初始化
	Init()

	// 自定义初始化
	for _, f := range inits {
		f()
	}

	config := global.Config.Server
	c := gin.Default()

	// 路由初始化
	initRouter(c, routers)

	err := c.Run(config.GetPort())
	if err != nil {
		log.Warn("start server error:%v", err)
		panic(err)
	}
}

func initRouter(r *gin.Engine, routers []types.RouterFunc) {
	for _, router := range routers {
		router(r)
	}
}
