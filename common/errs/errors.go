package errs

import "fmt"

//业务上的错误
var (
	ErrNoSuchUser             = &Error{Code: "ErrNoSuchUser", Message: "err no such user"}
	ErrNameOrPassword         = &Error{Code: "ErrNameOrPassword", Message: "err name or password"}
	ErrParam                  = &Error{Code: "ErrParam", Message: "err param"}
	ErrInternal               = &Error{Code: "ErrInternal", Message: "internal error"}
	ErrNotFound               = &Error{Code: "ErrNotFound", Message: "request not found"}
	ErrClientAccTokenExp      = &Error{Code: "ErrClientAccTokenExp", Message: "acc token is expiration"}
	ErrClientAccTokenNotValid = &Error{Code: "ErrClientAccTokenNotValid", Message: "token not valid"}
	UnAuthorized              = &Error{Code: "UnAuthorized", Message: "un authorized"}
)

//根据有的error创建一个新的error
func NewErr(e *Error, msg string) *Error {
	return &Error{Code: e.Code, Message: msg}
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}
