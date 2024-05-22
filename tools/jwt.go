package tools

import (
	"core/types"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Platform types.PlatformType // 平台
	Device   types.DeviceType   // 设备类型
	Token    string
}

const (
	TokenExpireDuration        = time.Minute * 1
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

var CustomSecret = []byte("夏天夏天悄悄过去")

func Token(id string, platform types.PlatformType, device types.DeviceType) (string, string, error) {

	token, err := GenToken(TokenExpireDuration, id, platform, device, nil)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenToken(RefreshTokenExpireDuration, id, platform, device, &token)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

// GenToken 生成JWT
func GenToken(expires time.Duration, id string, platform types.PlatformType, device types.DeviceType, str *string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
			Issuer:    "ohWin", // 签发人
			ID:        id,
		},
		Platform: platform,
		Device:   device,
	}
	if str != nil {

		claims.Token = *str
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

//func GenRefreshToken(id string, platform types.PlatformType, device types.DeviceType, str string) (string, error) {
//
//	claims := CustomClaims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireDuration)),
//			Issuer:    "ohWin", // 签发人
//		},
//		Random: randomStr(6),
//	}
//
//	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return refreshToken.SignedString(CustomSecret)
//}

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
			fmt.Println("Token is expired")
			return true
		}
	}
	return false
}
