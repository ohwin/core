package errorx

import (
	"core/types"
)

const (
	RespCodeTypeUnknown                     types.ErrCodeType = 100000 + iota // 未知错误
	RespCodeTypeServerInternal                                                // 服务器内部错误
	RespCodeTypeUserAlreadyExist                                              // 用户已存在
	RespCodeTypeTokenInvalid                                                  // Token无效
	RespCodeTypeTokenError                                                    // Token错误
	RespCodeTypeParameterError                                                // 参数错误
	RespCodeTypeIncorrectUsernameOrPassword                                   // 用户名或密码错误
)

var ErrMsgs = map[types.ErrCodeType]string{
	RespCodeTypeUnknown:                     "Unknown Error",
	RespCodeTypeServerInternal:              "Server Internal",                // 服务器内部错误
	RespCodeTypeUserAlreadyExist:            "User Already Exist",             // 用户已存在
	RespCodeTypeTokenInvalid:                "Token Invalid",                  // Token无效
	RespCodeTypeTokenError:                  "Token Error",                    // Token错误
	RespCodeTypeParameterError:              "Parameter Error",                // 参数错误
	RespCodeTypeIncorrectUsernameOrPassword: "Incorrect Username Or Password", // 用户名或密码错误
}

type ErrorCodeStr struct {
	Code types.ErrCodeType
	Msg  string
}

func (e *ErrorCodeStr) Error() string {
	return e.Msg
}

func new(code types.ErrCodeType) error {
	return &ErrorCodeStr{
		Code: code,
		Msg:  ErrMsg(code),
	}
}

func ErrMsg(code types.ErrCodeType) string {
	if msg, ok := ErrMsgs[code]; ok {
		return msg
	}
	return "Unknown error"
}

func Wrap(code types.ErrCodeType) error {
	return new(code)
}
