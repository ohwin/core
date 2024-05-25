package auth

import (
	"context"
	"github.com/gin-gonic/gin"
)

func Set(ctx *gin.Context, key string, value any) {
	ctx.Set(key, value)
}

func Get(ctx context.Context, key string) any {
	return ctx.(*gin.Context).Value(key)
}

func RemoteIP(ctx context.Context) string {
	c := ctx.(*gin.Context)
	return c.RemoteIP()
}

func UID(ctx context.Context) string {
	if uid := Get(ctx, "uid"); uid != nil {
		return uid.(string)
	}
	return ""
}
