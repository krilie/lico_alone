package cmodel

import "github.com/krilie/lico_alone/common/errs"

// 发生错误时的返回码
type ErrorReturn struct {
	Code    int     `json:"code" swaggo:"true,错误码" example:"400"`
	Message string  `json:"message" swaggo:"true,错误信息" example:"错误信息"`
	Detail  *string `json:"detail,omitempty" swaggo:"false,错误的详细信息，用于排查错误"  example:"错误的详细信息，用于排查错误"` // 可由运行模式控制是否显示
}

func NewRet(err *errs.Err) *ErrorReturn {
	detail := errs.GetErrMsg(err)
	return &ErrorReturn{
		Code:    err.Code,
		Message: err.Message,
		Detail:  &detail,
	}
}

func RetFromErr(err error) *ErrorReturn {
	Err := errs.GetInnerErr(err)
	if Err != nil {
		detail := errs.GetErrMsg(err)
		return &ErrorReturn{
			Code:    Err.Code,
			Message: Err.Message,
			Detail:  &detail,
		}
	} else {
		detail := err.Error()
		return &ErrorReturn{
			Code:    500,
			Message: "内部错误",
			Detail:  &detail,
		}
	}
}
