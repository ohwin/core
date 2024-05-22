package types

type ErrCodeType uint32

type LoginAndRegisterType uint8

const (
	LoginTypeNone LoginAndRegisterType = iota
	LoginTypeAccount
	LoginTypePhone
	LoginTypeEmail
)

//type RegisterType uint8

const (
	RegisterTypeNone LoginAndRegisterType = iota
	RegisterTypeAccount
	RegisterTypePhone
	RegisterTypeEmail
)

type PlatformType uint8

const (
	PlatformTypeNone PlatformType = iota // 所有平台
	PlatformTypeWeb
	PlatformTypeAndroid
	PlatformTypeIOS
)

type DeviceType uint8

const (
	DeviceTypeNone DeviceType = iota
	DeviceTypeWeb
	DeviceTypePhone
	DeviceTypePad
)

type BooleanType uint8

const (
	FirstLoginTypeNone BooleanType = iota
	FirstLoginTypeNo               // 今日首次登录
	FirstLoginTypeYes              // 不是今日首次登录
)
