package types

import "github.com/gin-gonic/gin"

type (
	StdStatusCode  uint8  // 标准状态类型
	StdBooleanType uint8  // 标准布尔类型
	ErrCodeType    uint32 // 错误类型
)

type (
	RouterFunc           func(engine *gin.Engine)
	LoginAndRegisterType uint8
	PlatformType         uint8
	DeviceType           uint8
)

const (
	StdStatusCodeNone  StdStatusCode = iota // 用户查询所有
	StdStatusCodeAllow                      // 表示可用、打开、有效
	StdStatusCodeDeny                       // 表示禁用、关闭、无效
)

const (
	LoginTypeNone LoginAndRegisterType = iota
	LoginTypeAccount
	LoginTypePhone
	LoginTypeEmail
)

const (
	RegisterTypeNone LoginAndRegisterType = iota
	RegisterTypeAccount
	RegisterTypePhone
	RegisterTypeEmail
)

const (
	PlatformTypeNone PlatformType = iota // 所有平台
	PlatformTypeWeb
	PlatformTypeAndroid
	PlatformTypeIOS
)

const (
	DeviceTypeNone DeviceType = iota
	DeviceTypeWeb
	DeviceTypePhone
	DeviceTypePad
)

const (
	FirstLoginTypeNone StdBooleanType = iota // 所有
	FirstLoginTypeNo                         // 今日首次登录
	FirstLoginTypeYes                        // 不是今日首次登录
)
