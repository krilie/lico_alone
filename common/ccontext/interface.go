package ccontext

import (
	"github.com/krilie/lico_alone/common/utils/id_util"
	"time"
)

func NewContext() *Context {
	ctx := &Context{}
	ctx.TraceId = id_util.GetUuid()
	ctx.StartTime = time.Now()
	return ctx
}

func MustGetContext(ctx interface{}) *Context {
	if c, ok := ctx.(*Context); !ok {
		panic("err convert from interface{} to *Context")
	} else {
		return c
	}
}

func GetContextOrNil(ctx interface{}) *Context {
	if c, ok := ctx.(*Context); !ok {
		return nil
	} else {
		return c
	}
}

func GetContextOrNew(ctx interface{}) *Context {
	if c, ok := ctx.(*Context); !ok {
		return NewContext()
	} else {
		return c
	}
}
