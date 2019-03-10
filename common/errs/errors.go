package errs

import "fmt"

var (
	ErrNoSuchUser     = &Error{Code: "ErrNoSuchUser", Message: "err no such user"}
	ErrNameOrPassword = &Error{Code: "ErrNameOrPassword", Message: "err name or password"}
	ErrParam          = &Error{Code: "ErrParam", Message: "err param"}
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
