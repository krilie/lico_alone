package errs

import "fmt"

//业务上的错误
var (
	ErrNoSuchUser     = &Error{Code: "ErrNoSuchUser", Message: "err no such user"}
	ErrNameOrPassword = &Error{Code: "ErrNameOrPassword", Message: "err name or password"}
	ErrParam          = &Error{Code: "ErrParam", Message: "err param"}
	ErrInternal       = &Error{Code: "ErrInternal", Message: "internal error"}
	ErrNotFound       = &Error{Code: "ErrNotFound", Message: "request not found"}
)

//根据有的error创建一个新的error
func NewErr(e *Error, msg string) *Error {
	return &Error{Code: e.Code, Message: msg}
}

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}
