package errs

import (
	"errors"
	"testing"
)

func TestGetErr(t *testing.T) {
	err := NewBadRequest().WithError(NewInternal().WithMsg("internal")).WithMsg("hello")
	t.Log(GetInnerErr(err))
	t.Log(GetErrMsg(err))
	t.Log(errors.Is(err, errBadRequest))
	t.Log(errors.As(err, &errBadRequest))
}
