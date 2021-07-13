package ginutil

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
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
	log    *nlog.NLog
	GinCtx *gin.Context
	AppCtx context2.Context
}

func NewGinWrap(ctx *gin.Context, log *nlog.NLog) *GinWrap {
	var wrap = &GinWrap{log: log, GinCtx: ctx}
	wrap.AppCtx = wrap.GetAppContext()
	wrap.log = wrap.log.Get(wrap.AppCtx, "gin wrap")
	return wrap
}

func (c *GinWrap) BindQuery(out interface{}) {
	err := c.GinCtx.ShouldBindQuery(out)
	if err != nil {
		panic(errs.NewParamError().WithMsg(err.Error()))
	}
}
func (c *GinWrap) BindForm(out interface{}) {
	err := c.GinCtx.ShouldBind(out)
	if err != nil {
		panic(errs.NewParamError().WithMsg(err.Error()))
	}
}
func (c *GinWrap) BindBodyJson(out interface{}) {
	err := c.GinCtx.BindJSON(out)
	if err != nil {
		panic(errs.NewParamError().WithMsg(err.Error()))
	}
}
func (c *GinWrap) BindPath(out interface{}) {
	err := c.GinCtx.BindUri(out)
	if err != nil {
		panic(errs.NewParamError().WithMsg(err.Error()))
	}
}
func (c *GinWrap) BindHeader(out interface{}) {
	err := c.GinCtx.BindHeader(out)
	if err != nil {
		panic(errs.NewParamError().WithMsg(err.Error()))
	}
}
