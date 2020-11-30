package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/nlog"
)

// some const value for gin http protocol
var (
	HeaderClientAccToken = "ClientAccToken" // for header client access token
	HeaderTraceId        = "TraceId"        //for header trace id
	HeaderAuthorization  = "Authorization"  //for authorization
	GinKeyAppContext     = "context"
)

type GinUtils struct {
	log *nlog.NLog
}

func NewGinUtils(log *nlog.NLog) *GinUtils {
	return &GinUtils{log: log}
}

// 给中间件使用
// get app context or nil
func (g *GinUtils) GetAppCtxOrAbort(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		g.log.WithFuncName("GetAppCtxOrAbort").Error("can not get service context for next step")
		c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	contextOrNil := context.GetContextOrNil(value)
	if contextOrNil == nil {
		g.log.WithFuncName("GetAppCtxOrAbort").Error("GetAppCtxOrAbort", "internal err on cast context to app context")
		c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	return contextOrNil
}

func (g *GinUtils) GetUserIdOrAbort(c *gin.Context) string {
	ctx := g.GetAppCtxOrAbort(c)
	if ctx == nil {
		return ""
	}
	if ctx.GetUserId() == "" {
		AbortWithErr(c, errs.NewInvalidToken().WithMsg("登录信息无效"))
		return ""
	} else {
		return ctx.GetUserId()
	}
}

// GetAppCtxOrReturn get app context or nil
func (g *GinUtils) GetAppCtxOrReturn(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		g.log.WithFuncName("GetAppCtxOrAbort").Error("GetAppCtxOrReturn", "can not get service context for next step")
		c.JSON(200, com_model.NewRet(errs.NewInternal().WithMsg("ctx not get")))
		return nil
	}
	contextOrNil := context.GetContextOrNil(value)
	if contextOrNil == nil {
		g.log.WithFuncName("GetAppCtxOrAbort").Error("GetAppCtxOrReturn", "internal err on cast context to app context")
		c.JSON(200, com_model.NewRet(errs.NewInternal().WithMsg("ctx not get is nil")))
		return nil
	}
	return contextOrNil
}

func (g *GinUtils) MustGetAppCtx(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		g.log.WithFuncName("GetAppCtxOrAbort").Panic("GetAppCtxOrReturn", "can not get service context for next step")
		return nil
	}
	return context.MustGetContext(value)
}

func (g *GinUtils) MustGetUserId(c *gin.Context) string {
	ctx := g.MustGetAppCtx(c)
	if ctx.GetUserId() != "" {
		return ctx.GetUserId()
	} else {
		g.log.WithFuncName("GetAppCtxOrAbort").Panic("must get user id can not get user id.")
		return ""
	}
}
