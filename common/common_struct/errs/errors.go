package errs

import (
	"fmt"
	"github.com/lico603/lico_user/common/common_struct"
)

//业务上的错误
var (
	ErrNoSuchUser             = &Error{HttpStatus: 404, Code: "ErrNoSuchUser", Message: "err no such user"}
	ErrNameOrPassword         = &Error{HttpStatus: 400, Code: "ErrNameOrPassword", Message: "err name or password"}
	ErrParam                  = &Error{HttpStatus: 400, Code: "ErrParam", Message: "err param"}
	ErrInternal               = &Error{HttpStatus: 500, Code: "ErrInternal", Message: "internal error"}
	ErrNotFound               = &Error{HttpStatus: 404, Code: "ErrNotFound", Message: "request not found"}
	ErrClientAccTokenExp      = &Error{HttpStatus: 401, Code: "ErrClientAccTokenExp", Message: "acc token is expiration"}
	ErrClientAccTokenNotValid = &Error{HttpStatus: 401, Code: "ErrClientAccTokenNotValid", Message: "token not valid"}
	UnAuthorized              = &Error{HttpStatus: 401, Code: "UnAuthorized", Message: "un authorized"}
	ErrNoPermission           = &Error{HttpStatus: 403, Code: "ErrNoPermission", Message: "no permission"}
)

type Error struct {
	HttpStatus int    //ref http status
	Code       string //code of error
	Message    string //message of error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}

func (e *Error) NewWithMsg(msg string) *Error {
	return &Error{HttpStatus: e.HttpStatus, Code: e.Code, Message: msg}
}

func (e *Error) NewAppendMsg(msg string) *Error {
	return &Error{HttpStatus: e.HttpStatus, Code: e.Code, Message: e.Message + ":" + msg}
}

func (e *Error) ToStdReturn() *common_struct.StdReturn {
	return &common_struct.StdReturn{Code: e.Code, Message: e.Message}
}
func (e *Error) ToStdWithMsg(msg string) *common_struct.StdReturn {
	return &common_struct.StdReturn{Code: e.Code, Message: msg}
}
func (e *Error) ToStdAppendMsg(msg string) *common_struct.StdReturn {
	return &common_struct.StdReturn{Code: e.Code, Message: e.Message + ":" + msg}
}
