package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ohwin/core/errorx"
	"github.com/ohwin/core/htp"
	"github.com/ohwin/core/tools"
	"github.com/ohwin/core/types"
	"strings"
)

func JWTAuth(skipRouters []types.SkipRouter) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.Path
		method := ctx.Request.Method
		for _, skip := range skipRouters {
			if skip.Method == method && skip.Url == url {
				ctx.Next()
				return
			}
		}

		// 获取token
		token, tokenErr := getToken(ctx, htp.Authorization)
		refreshToken, refreshTokenErr := getToken(ctx, htp.RefreshToken)
		if tokenErr != nil || refreshTokenErr != nil {
			htp.Fail(ctx, errorx.RespCodeTypeTokenError).Abort()
			return
		}

		// 解析token
		_, tokenErr = tools.ParseToken(token)
		refreshClaims, refreshTokenErr := tools.ParseToken(refreshToken)

		// refreshToken错误或者token与refreshToken中保存的token不一致直接返回
		if refreshTokenErr != nil || !tools.TokenCompare(token, refreshClaims.Token) {
			htp.Fail(ctx, errorx.RespCodeTypeTokenError).Abort()
			return
		}

		if tokenErr != nil {
			if tools.TokenIsExpiredErr(tokenErr) { // 过期错误刷新Token
				newToken, newRefreshToken, err := tools.Token(refreshClaims.ID, refreshClaims.Platform, refreshClaims.Device, refreshClaims.Role)
				if err != nil {
					htp.Fail(ctx, errorx.RespCodeTypeServerInternal).Abort()
					return
				}
				htp.SetHeader(ctx, htp.Authorization, newToken)
				htp.SetHeader(ctx, htp.RefreshToken, newRefreshToken)
			} else { // token无效
				htp.Fail(ctx, errorx.RespCodeTypeTokenInvalid).Abort()
				return
			}
		}

		// 平台验证 Todo

		// token有效
		if tools.TokenCompare(token, refreshClaims.Token) {
			ctx.Set("uid", tools.StrToUint(refreshClaims.ID))
			ctx.Next()
		}
	}
}

func getToken(ctx *gin.Context, header string) (string, error) {
	authHeader := ctx.Request.Header.Get(header)
	if authHeader == "" {
		return "", errorx.Wrap(errorx.RespCodeTypeTokenError)
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errorx.Wrap(errorx.RespCodeTypeTokenError)
	}

	return parts[1], nil
}
