package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/log"
)

func CORS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//设置响应头中的"Access-Control-Allow-Origin"字段，允许任何域名访问资源（这是一个简单的CORS策略，实际应用中可能需要更严格的控制）。
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Authorization,Refresh-Token, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin,Refresh-Token, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
		ctx.Header("Access-Control-Max-Age", "172800")                                                                                                                                                                         // 缓存请求信息 单位为秒
		ctx.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                                //  跨域请求是否需要带cookie信息 默认设置为true
		ctx.Set("content-type", "application/yaml")                                                                                                                                                                            // 设置返回格式是yaml

		//放行所有OPTIONS方法
		if ctx.Request.Method == "OPTIONS" {
			log.Debug("options")
			ctx.AbortWithStatus(200)
		}

	}
}
