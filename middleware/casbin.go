package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/auth"
	"github.com/ohwin/core/errorx"
	"github.com/ohwin/core/global"
	"github.com/ohwin/core/htp"
	"github.com/ohwin/core/types"
	"log"
)

func Casbin(skipRouters []types.SkipRouter) gin.HandlerFunc {
	e := global.Enforce
	return func(ctx *gin.Context) {

		url := ctx.Request.URL.Path
		method := ctx.Request.Method
		for _, skip := range skipRouters {
			if skip.Method == method && skip.Url == url {
				ctx.Next()
				return
			}
		}

		err := e.LoadPolicy()
		if err != nil {
			htp.Fail(ctx, errorx.RespCodeTypeTokenError).Abort()
			return
		}
		// 简便起见，假设用户从url传递 /xxxx?username=leo，实际应用可以结合jwt等鉴权
		uid := auth.UID(ctx)
		log.Println(uid, ctx.Request.URL.Path, ctx.Request.Method)
		ok, err := e.Enforce(uid, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			htp.Fail(ctx, errorx.RespCodeTypeTokenError).Abort()
			return
		} else if !ok {
			htp.Fail(ctx, errorx.RespCodeTypeTokenError).Abort()
			return
		}
		ctx.Next()
	}
}
