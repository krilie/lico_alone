package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/model"
	"github.com/krilie/lico_alone/component/nlog"
)

// some const value for gin http protocol
var (
	HeaderClientAccToken = "ClientAccToken" // for header client access token
	HeaderTraceId        = "TraceId"        //for header trace id
	HeaderAuthorization  = "Authorization"  //for authorization
	GinKeyAppContext     = "context"
)

// 给中间件使用
// get app context or nil
func GetAppCtxOrAbort(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		nlog.Error("GetAppCtxOrAbort", "can not get application context for next step")
		c.AbortWithStatusJSON(500, model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	contextOrNil := context.GetContextOrNil(value)
	if contextOrNil == nil {
		nlog.Error("GetAppCtxOrAbort", "internal err on cast context to app context")
		c.AbortWithStatusJSON(500, model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	return contextOrNil
}

func GetUserIdOrAbort(c *gin.Context) string {
	ctx := GetAppCtxOrAbort(c)
	if ctx == nil {
		return ""
	}
	if ctx.GetUserId() == "" {
		AbortWithErr(c, errs.NewUnauthorized().WithMsg("没取到用户id"))
		return ""
	} else {
		return ctx.GetUserId()
	}
}

// GetAppCtxOrReturn get app context or nil
func GetAppCtxOrReturn(c *gin.Context) *context.Context {
	nlog = nlog.WithField(nlog.Function, "GetAppCtxOrReturn")
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		nlog.Error("GetAppCtxOrReturn", "can not get application context for next step")
		c.JSON(500, model.NewRet(errs.NewInternal().WithMsg("ctx not get")))
		return nil
	}
	contextOrNil := context.GetContextOrNil(value)
	if contextOrNil == nil {
		nlog.Error("GetAppCtxOrReturn", "internal err on cast context to app context")
		c.JSON(500, model.NewRet(errs.NewInternal().WithMsg("ctx not get is nil")))
		return nil
	}
	return contextOrNil
}

func MustGetAppCtx(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		nlog.Panic("GetAppCtxOrReturn", "can not get application context for next step")
		return nil
	}
	return context.MustGetContext(value)
}

func MustGetUserId(c *gin.Context) string {
	ctx := MustGetAppCtx(c)
	if ctx.GetUserId() != "" {
		return ctx.GetUserId()
	} else {
		nlog.Panic("must get user id can not get user id.")
		return ""
	}
}
