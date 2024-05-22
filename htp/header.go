package htp

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HeaderType map[string]any

const (
	Authorization = "Authorization"
	RefreshToken  = "Refresh-Token"
)

func SetHeader(ctx *gin.Context, k string, v any) {
	switch k {
	case Authorization:
		ctx.Header(k, fmt.Sprintf("Bearer %s", v.(string)))
	case RefreshToken:
		ctx.Header(k, fmt.Sprintf("Bearer %s", v.(string)))
	default:
		ctx.Header(k, v.(string))
	}
}
