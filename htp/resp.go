package htp

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/errorx"
	"github.com/ohwin/core/types"
)

type resp struct {
	Msg  string            `yaml:"msg"`
	Code types.ErrCodeType `yaml:"code"`
	Data any               `yaml:"data,omitempty"`
}

func OK(ctx *gin.Context, data any) {
	ctx.JSON(200, &resp{
		Msg:  "success",
		Code: 200,
		Data: data,
	})
}

func OKWithMsg(ctx *gin.Context, data any, msg string) {
	ctx.JSON(200, &resp{
		Msg:  msg,
		Code: 200,
		Data: data,
	})
}

func OKWithHeader(ctx *gin.Context, data any, headers map[string]any) {
	for k, v := range headers {
		SetHeader(ctx, k, v)
	}
	ctx.JSON(200, &resp{
		Msg:  "success",
		Code: 200,
		Data: data,
	})
}

func Fail(ctx *gin.Context, code types.ErrCodeType) {
	ctx.JSON(200, &resp{
		Msg:  errorx.ErrMsg(code),
		Code: code,
	})
}

func FailWithErr(ctx *gin.Context, err error) {
	errorCodeStr := err.(*errorx.ErrorCodeStr)
	ctx.JSON(200, &resp{
		Msg:  errorCodeStr.Msg,
		Code: errorCodeStr.Code,
	})

}
