package ginutil

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/component/nlog"
)

// some const value for gin http protocol
var (
	HeaderClientAccToken = "ClientAccToken" // for header client access token
	HeaderTraceId        = "TraceId"        //for header trace id
	HeaderAuthorization  = "Authorization"  //for authorization
	GinKeyAppContext     = "GinKeyAppContext"
)

type GinWrap struct {
	log *nlog.NLog
	*gin.Context
	AppCtx context2.Context
}

func NewGinWrap(ctx *gin.Context, log *nlog.NLog) *GinWrap {
	var wrap = &GinWrap{log: log, Context: ctx}
	wrap.AppCtx = wrap.GetAppContext()
	wrap.log = wrap.log.Get(wrap.AppCtx, "gin wrap")
	return wrap
}
