package errs

import (
	"errors"
	"testing"
)

func TestGetErr(t *testing.T) {
	err := NewNormal().WithError(NewNormal().WithMsg("internal")).WithMsg("hello")
	t.Log(GetInnerErr(err))
	t.Log(GetErrMsg(err))
	t.Log(errors.Is(err, InternalError))
	t.Log(errors.As(err, &NormalError))
}
