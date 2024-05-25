package core

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/initialize"
)

func Init() {
	initialize.Viper()
	initialize.Log()
	initialize.Redis()
	//initialize.MQ()
	initialize.DB()
}

func RunWindowsServer() {

	c := gin.Default()

	{
		initialize.Routers(c)
	}

	c.Run(":8080")
}
