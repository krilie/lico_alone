package common_struct

import "github.com/lico603/lico-my-site-user/common/common_struct/errs"

//常规返回类型
var (
	StdSuccess = &StdReturn{Code: "success", Message: "success"}
	StdFailure = &StdReturn{Code: "failure", Message: "failure"}
)

// std return type for http
type StdReturn struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (s *StdReturn) WithCode(code string) {
	s.Code = code
}

func (s *StdReturn) WithMessage(msg string) {
	s.Message = msg
}

func NewStdRet(code, msg string) *StdReturn {
	return &StdReturn{Code: code, Message: msg}
}
func NewStdRetFromErr(e *errs.Error) *StdReturn {
	return &StdReturn{e.Code, e.Message}
}
func NewStdRetFromErrWithMsg(e *errs.Error, msg string) *StdReturn {
	return &StdReturn{e.Code, msg}
}
func NewStdRetFromErrAppendMsg(e *errs.Error, msg string) *StdReturn {
	return &StdReturn{e.Code, e.Message + ":" + msg}
}
