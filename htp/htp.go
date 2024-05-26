package htp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/errorx"
	"github.com/ohwin/core/types"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

const (
	Authorization = "Authorization"
	RefreshToken  = "Refresh-Token"
)

func SetHeader(ctx *gin.Context, k string, v string) {
	switch k {
	case Authorization:
		ctx.Header(k, fmt.Sprintf("Bearer %s", v))
	case RefreshToken:
		ctx.Header(k, fmt.Sprintf("Bearer %s", v))
	default:
		ctx.Header(k, v)
	}
}

type resp struct {
	Msg  string            `json:"msg"`
	Code types.ErrCodeType `json:"code"`
	Data any               `json:"data"`
}

// OK 成功响应
func OK(ctx *gin.Context, data any) {
	ctx.JSON(200, &resp{
		Msg:  "success",
		Code: 200,
		Data: data,
	})
}

// OKWithMsg 成功响应携带自定义消息
func OKWithMsg(ctx *gin.Context, data any, msg string) {
	ctx.JSON(200, &resp{
		Msg:  msg,
		Code: 200,
		Data: data,
	})
}

// OKWithHeader 成功响应携带自定义请求头
func OKWithHeader(ctx *gin.Context, data any, headers map[string]string) {
	for k, v := range headers {
		SetHeader(ctx, k, v)
	}
	ctx.JSON(200, &resp{
		Msg:  "success",
		Code: 200,
		Data: data,
	})
}

// Fail 失败响应
func Fail(ctx *gin.Context, code types.ErrCodeType) *gin.Context {
	ctx.JSON(200, &resp{
		Msg:  errorx.ErrMsg(code),
		Code: code,
	})
	return ctx
}

// FailWithErr 失败响应携带自定义错误
func FailWithErr(ctx *gin.Context, err error) *gin.Context {
	errorCodeStr := err.(*errorx.ErrorCodeStr)
	ctx.JSON(200, &resp{
		Msg:  errorCodeStr.Msg,
		Code: errorCodeStr.Code,
	})
	return ctx
}
