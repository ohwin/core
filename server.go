package core

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/initialize"
)

func RunWindowsServer() {
	initialize.Viper()
	initialize.Log()
	initialize.Redis()
	//initialize.MQ()
	initialize.DB()

	c := gin.Default()

	{
		initialize.Routers(c)
	}

	c.Run(":8080")
}
