package initialize

import (
	"core/middleware"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	//r = r.Group("", middleware.CORS())
	r.Use(middleware.CORS())
	//routerPublic := r.Group("")
	//routerProtect := r.Group("", middleware.JWTAuth())
	//ro.InitBaseRouter(routerPublic)
	//router.InitUserRouter(routerProtect)
	//router.InitQuestionRouter(routerPublic)

}
