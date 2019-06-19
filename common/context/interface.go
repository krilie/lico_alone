package context

import (
	"github.com/krilie/lico_alone/common/utils/id_util"
	"time"
)

type Context interface {
	GetTraceId() string
	GetStartTime() time.Time
	GetLastTime() time.Time
	GetClientId() string
	GetClientToken() string
	GetUserId() string
	GetUserToken() string
	GetUserNickName() string
	GetUserLoginName() string
	SetTraceId(string)
	SetStartTime(time.Time)
	SetLastTime(time.Time)
	SetClientId(string)
	SetClientToken(string)
	SetUserId(string)
	SetUserToken(string)
	SetUserNickName(string)
	SetUserLoginName(string)
}

func NewContext() Context {
	ctx := &context{}
	ctx.TraceId = id_util.GetUuid()
	ctx.StartTime = time.Now()
	return ctx
}

func MustGetContext(ctx interface{}) Context {
	if c, ok := ctx.(Context); !ok {
		panic("err convert from interface{} to *context")
	} else {
		return c
	}
}

func GetContextOrNil(ctx interface{}) Context {
	if c, ok := ctx.(Context); !ok {
		return nil
	} else {
		return c
	}
}
