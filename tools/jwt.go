package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ohwin/core/log"
	"github.com/ohwin/core/types"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Platform types.PlatformType // 平台
	Device   types.DeviceType   // 设备类型
	UID      string             // 用户UID
	Role     string             // 用户角色，逗号分隔
	Token    string
}

const (
	TokenExpireDuration        = time.Minute * 1
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

var CustomSecret = []byte("夏天夏天悄悄过去")

func Token(id string, platform types.PlatformType, device types.DeviceType, role string) (string, string, error) {

	token, err := GenToken(TokenExpireDuration, platform, device, id, "", role)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenToken(RefreshTokenExpireDuration, platform, device, token, id, role)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

// GenToken 生成JWT
func GenToken(expires time.Duration, platform types.PlatformType, device types.DeviceType, id, tokenStr, role string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
			Issuer:    "ohWin", // 签发人
			ID:        id,
		},
		Platform: platform,
		Device:   device,
		UID:      id,
		Role:     role,
		Token:    tokenStr,
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func TokenCompare(token1, token2 string) bool {
	return token1 == token2
}

// TokenIsExpiredErr 通过解析Token返回的错误判断是否为过期错误
// err为过期错误返回true，为空或其他错误均返回false
func TokenIsExpiredErr(err error) bool {
	if validationErr, ok := err.(*jwt.ValidationError); ok {
		if validationErr.Errors == jwt.ValidationErrorExpired {
			log.Debug("Token Error %s", "Token is expired")
			return true
		}
	}
	return false
}
