package cmodel

import (
	"github.com/krilie/lico_alone/common/errs"
)

var StdSuccess = &CommonReturn{
	Code:    200,
	Message: "success",
	Detail:  nil,
	Data:    nil,
}

type CommonReturn struct {
	Code    int         `json:"code" swaggo:"true,错误码" example:"400"`
	Message string      `json:"message" swaggo:"true,错误信息" example:"错误信息"`
	Detail  *string     `json:"detail,omitempty" swaggo:"false,错误的详细信息，用于排查错误"  example:"错误的详细信息，用于排查错误"` // 可由运行模式控制是否显示
	Data    interface{} `json:"data,omitempty" `                                                          // 数据值
}

func NewRet(err *errs.Err) *CommonReturn {
	fullMsg := err.GetFullMsg()
	return &CommonReturn{
		Code:    err.Code,
		Message: err.Message,
		Detail:  &fullMsg,
		Data:    nil,
	}
}
func NewRetFromErr(err error) *CommonReturn {
	fullMsg := errs.GetErrMsg(err)
	return &CommonReturn{
		Code:    500,
		Message: err.Error(),
		Detail:  &fullMsg,
		Data:    nil,
	}
}
func NewFailure(code int, msg string) *CommonReturn {
	if msg == "" {
		msg = "failure"
	}
	return &CommonReturn{
		Code:    code,
		Message: msg,
		Detail:  nil,
		Data:    nil,
	}
}
