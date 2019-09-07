package errs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 定义通用错误 pkg errors cause error
var (
	Error           = &Err{Code: 0, Message: ""}        // nil
	ErrForbidden    = &Err{Code: 403, Message: "禁止访问"}  // 403
	ErrNotFound     = &Err{Code: 404, Message: "资源不存在"} // 404
	ErrBadRequest   = &Err{Code: 400, Message: "请求无效"}  // 400
	ErrUnauthorized = &Err{Code: 401, Message: "未授权"}   // 401
	ErrInternal     = &Err{Code: 500, Message: "服务器错误"} // 500
	// 其它错误
	ErrDbCreate = &Err{Code: 500, Message: "数据库创建错误"}
	ErrDbUpdate = &Err{Code: 500, Message: "数据库更新错误"}
	ErrDbDelete = &Err{Code: 500, Message: "数据库删除错误"}
	ErrDbQuery  = &Err{Code: 500, Message: "数据库查询错误"}
)

type Err struct {
	Code    int
	Message string
	Err     error // 原始错误
}

func (w *Err) Error() string {
	builder := strings.Builder{}
	builder.WriteString("[code:")
	builder.WriteString(strconv.Itoa(w.Code))
	builder.WriteString(" message:")
	builder.WriteString(w.Message)
	if w.Err != nil {
		builder.WriteString(" err:")
		builder.WriteString(w.Err.Error())
		builder.WriteString("]")
	}
	return builder.String()
}

func (w *Err) WithError(err error) *Err { w.Err = err; return w }
func (w *Err) GetCode() int             { return w.Code }
func (w *Err) WithCode(code int) *Err   { w.Code = code; return w }
func (w *Err) WithMsg(msg string) *Err  { w.Message = msg; return w }
func (w *Err) New() *Err                { return &Err{Code: w.Code, Message: w.Message, Err: w.Err} }
func (w *Err) Unwrap() error            { return w.Err }

func (w *Err) WithMsgf(format string, args ...interface{}) *Err {
	w.Message = fmt.Sprintf(format, args...)
	return w
}

func (w *Err) GetFullMsg() string {
	return GetErrMsg(w)
}

func New() *Err             { return &Err{} }
func NewForbidden() *Err    { return ErrForbidden.New() }
func NewNotFound() *Err     { return ErrNotFound.New() }
func NewBadRequest() *Err   { return ErrBadRequest.New() }
func NewUnauthorized() *Err { return ErrUnauthorized.New() }
func NewInternal() *Err     { return ErrInternal.New() }

// 其它错误
func NewErrDbCreate() *Err { return ErrDbCreate.New() }
func NewErrDbUpdate() *Err { return ErrDbUpdate.New() }
func NewErrDbDelete() *Err { return ErrDbDelete.New() }
func NewErrDbQuery() *Err  { return ErrDbQuery.New() }

// 取到最内层的Err 如是没有返回nil
func GetInnerErr(err error) *Err {
	var retErr *Err
	for err != nil {
		tErr, ok := err.(*Err)
		if ok {
			retErr = tErr
			err = tErr.Err
			continue
		} else {
			break
		}
	}
	return retErr
}

func ToErrOrNil(err error) *Err {
	if err == nil {
		return nil
	}
	tErr, ok := err.(*Err)
	if ok {
		return tErr
	}
	return nil
}

func GetCode(err error) int {
	Err := ToErrOrNil(err)
	if Err != nil {
		return Err.Code
	}
	return 500
}

// 循环取出所有错误信息
func GetErrMsg(err error) string {
	var strBuild strings.Builder
	for err != nil {
		tErr, ok := err.(*Err)
		if ok {
			strBuild.WriteString(tErr.Message)
			strBuild.WriteString("|")
			err = errors.Unwrap(err)
			continue
		} else {
			strBuild.WriteString(err.Error())
			break
		}
	}
	return strBuild.String()
}
