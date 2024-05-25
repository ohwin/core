package core

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/global"
	"github.com/ohwin/core/initialize"
	"github.com/ohwin/core/log"
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

func RunWindowsServer() {
	Init()

	config := global.Config.Server
	c := gin.Default()
	{
		initialize.Routers(c)
	}

	err := c.Run(config.Port)
	if err != nil {
		log.Warn("start server error:%v", err)
		panic(err)
	}
}
