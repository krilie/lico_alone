package errs

import (
	"fmt"
)

//业务上的错误
var (
	ErrNoSuchUser             = &Error{RefHttpStatus: 404, Code: "ErrNoSuchUser", Message: "err no such user"}
	ErrNameOrPassword         = &Error{RefHttpStatus: 400, Code: "ErrNameOrPassword", Message: "err name or password"}
	ErrParam                  = &Error{RefHttpStatus: 400, Code: "ErrParam", Message: "err param"}
	ErrInternal               = &Error{RefHttpStatus: 500, Code: "ErrInternal", Message: "internal error"}
	ErrNotFound               = &Error{RefHttpStatus: 404, Code: "ErrNotFound", Message: "request not found"}
	ErrClientAccTokenExp      = &Error{RefHttpStatus: 401, Code: "ErrClientAccTokenExp", Message: "acc token is expiration"}
	ErrClientAccTokenNotValid = &Error{RefHttpStatus: 401, Code: "ErrClientAccTokenNotValid", Message: "token not valid"}
	UnAuthorized              = &Error{RefHttpStatus: 401, Code: "UnAuthorized", Message: "un authorized"}
)

type Error struct {
	RefHttpStatus int    //ref http status
	Code          string //code of error
	Message       string //message of error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}

func (e *Error) NewWithMsg(msg string) *Error {
	return &Error{RefHttpStatus: e.RefHttpStatus, Code: e.Code, Message: msg}
}

func (e *Error) NewAppendMsg(msg string) *Error {
	return &Error{RefHttpStatus: e.RefHttpStatus, Code: e.Code, Message: e.Message + ":" + msg}
}

func (e *Error) ToStdReturn() *StdReturn {
	return &StdReturn{Code: e.Code, Message: e.Message}
}
func (e *Error) ToStdWithMsg(msg string) *StdReturn {
	return &StdReturn{Code: e.Code, Message: msg}
}
func (e *Error) ToStdAppendMsg(msg string) *StdReturn {
	return &StdReturn{Code: e.Code, Message: e.Message + ":" + msg}
}

// std return type for http
type StdReturn struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
