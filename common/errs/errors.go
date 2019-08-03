package errs

import (
	"fmt"
	"strconv"
	"strings"
)

// 定义通用错误 pkg errors cause error
var (
	ErrForbidden    = &Err{Code: 403, Message: "禁止访问"}  // 403
	ErrNotFound     = &Err{Code: 404, Message: "资源不存在"} // 404
	ErrBadRequest   = &Err{Code: 400, Message: "请求无效"}  // 400
	ErrUnauthorized = &Err{Code: 401, Message: "未授权"}   // 401
	ErrInternal     = &Err{Code: 500, Message: "服务器错误"} // 500
)

type Err struct {
	Code    int
	Message string
	Err     error // 原始错误
}

func (e *Err) Error() string {
	builder := strings.Builder{}
	builder.WriteString("[code:")
	builder.WriteString(strconv.Itoa(e.Code))
	builder.WriteString(" message:")
	builder.WriteString(e.Message)
	if e.Err != nil {
		builder.WriteString(" err:")
		builder.WriteString(e.Err.Error())
		builder.WriteString("]")
	}
	return builder.String()
}
func (w *Err) Cause() error             { return w.Err }
func (w *Err) WithError(err error) *Err { w.Err = err; return w }
func (w *Err) GetCode() int             { return w.Code }
func (w *Err) WithCode(code int) *Err   { w.Code = code; return w }
func (w *Err) WithMsg(msg string) *Err  { w.Message = msg; return w }
func (e *Err) New() *Err                { return &Err{Code: e.Code, Message: e.Message, Err: e.Err} }

func (w *Err) WithMsgf(format string, args ...interface{}) *Err {
	w.Message = fmt.Sprintf(format, args...)
	return w
}

func New() *Err          { return &Err{} }
func Forbidden() *Err    { return ErrForbidden.New() }
func NotFound() *Err     { return ErrNotFound.New() }
func BadRequest() *Err   { return ErrBadRequest.New() }
func Unauthorized() *Err { return ErrUnauthorized.New() }
func Internal() *Err     { return ErrInternal.New() }

func GetErr(err error) *Err {
	type causer interface {
		Cause() error
	}
	for err != nil {
		tErr, ok := err.(*Err)
		if ok {
			return tErr
		}
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return nil
}
