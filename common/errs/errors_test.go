package errs

import (
	"github.com/pkg/errors"
	"testing"
)

func TestErr(t *testing.T) {
	err := errors.WithMessage(ErrInternal, "数据库插入错误")
	if ErrInternal == errors.Cause(err) {
		t.Log("ok")
	}
}

func TestErr2(t *testing.T) {
	err := errors.WithMessage(ErrBadRequest, "数据库插入错误")
	if ErrInternal == errors.Cause(err) {
		t.Log("ok")
	}
}

func TestErr3(t *testing.T) {
	err := errors.WithMessage(errors.New("fasdfas"), "数据库插入错误")
	err = errors.WithMessage(err, "数据库插入错误")
	err = errors.WithMessage(err, "数据库插入错误")
	err = errors.WithMessage(err, "数据库插入错误")
	err = errors.WithStack(err)
	if caser := errors.Cause(err); caser != nil {
		t.Log(caser.Error())
	}
	getErr := GetInnerErr(err)
	if getErr != nil {
		t.Log(getErr.GetCode())
		t.Log(getErr.Message)
	}
}

func TestUse1(t *testing.T) {
	// 构建错误对象
	err := NewInternal().WithMsg("这是个内部错误").WithError(errors.New("这是最内层的causer"))
	// 取到最内层错误
	if caser := errors.Cause(err); caser != nil {
		t.Log(caser.Error()) // 这是最内层的causer
	}
	// 取到自定义数据类型
	getErr := GetInnerErr(err)
	if getErr != nil {
		t.Log(getErr.GetCode()) // 500
		t.Log(getErr.Message)   // 这是个内部错误
	}
}

func TestUse2(t *testing.T) {
	// 构建错误对象
	err := NewInternal().WithMsg("这是个内部错误").WithError(errors.New("这是最内层的causer"))
	err2 := errors.Wrap(err, "这是个wrap")
	err2 = errors.WithMessage(err2, "这是个wrap")
	t.Logf("%v", err2)
	// 取到最内层错误
	if caser := errors.Cause(err2); caser != nil {
		t.Log(caser.Error()) // 这是最内层的causer
	}
	// 取到自定义数据类型
	getErr := GetInnerErr(err2)
	if getErr != nil {
		t.Log(getErr.GetCode()) // 500
		t.Log(getErr.Message)   // 这是个内部错误
	}
	// 输出信息栈

}
