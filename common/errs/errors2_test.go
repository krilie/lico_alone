package errs

import (
	"errors"
	"testing"
)

func TestGetErr(t *testing.T) {
	err := NewBadRequest().WithError(NewInternal().WithMsg("internal")).WithMsg("hello")
	t.Log(GetInnerErr(err))
	t.Log(GetErrMsg(err))
	t.Log(errors.Is(err, ErrBadRequest))
	t.Log(errors.As(err, &ErrBadRequest))
}
