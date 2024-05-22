package core

import (
	"core/initialize"
	"github.com/gin-gonic/gin"
)

func RunWindowsServer() {
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
