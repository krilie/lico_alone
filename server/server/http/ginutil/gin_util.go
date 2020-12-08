package ginutil

import (
	context2 "context"
	"errors"
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
	GinKeyAppContext     = "GinKeyAppContext"
)

type GinUtils struct {
	log *nlog.NLog
}

func NewGinUtils(log *nlog.NLog) *GinUtils {
	return &GinUtils{log: log}
}

// 给中间件使用
// get app context or nil
func (g *GinUtils) GetAppValuesOrAbort(c *gin.Context) *context.AppCtxValues {
	appContext := g.GetAppContext(c)
	if appContext == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("can not get service context for next step")
		c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	value := context.GetAppValues(appContext)
	if value == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("can not get service context for next step")
		c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	return value
}

func (g *GinUtils) GetAppContextOrAbort(c *gin.Context) context2.Context {
	appContext := g.GetAppContext(c)
	if appContext == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("can not get service context for next step")
		c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	return appContext
}

func (g *GinUtils) GetUserIdOrAbort(c *gin.Context) string {
	values := g.GetAppValuesOrAbort(c)
	if values == nil {
		return ""
	}
	if values.UserId == "" {
		AbortWithErr(c, errs.NewInvalidToken().WithMsg("登录信息无效"))
		return ""
	} else {
		return values.UserId
	}
}

// GetAppValuesOrReturn get app context or nil
func (g *GinUtils) GetAppValuesOrReturn(c *gin.Context) *context.AppCtxValues {
	appContext := g.GetAppContext(c)
	if appContext == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("GetAppValuesOrReturn", "can not get service context for next step")
		c.JSON(200, com_model.NewRet(errs.NewInternal().WithMsg("ctx not get")))
		return nil
	}
	value := context.GetAppValues(appContext)
	if value == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("GetAppValuesOrReturn", "can not get service context for next step")
		c.JSON(200, com_model.NewRet(errs.NewInternal().WithMsg("ctx not get")))
		return nil
	}
	return value
}

func (g *GinUtils) MustGetAppValues(c *gin.Context) *context.AppCtxValues {
	appContext := g.MustGetAppContext(c)
	value := context.MustGetAppValues(appContext)
	return value
}

func (g *GinUtils) GetUserId(c *gin.Context) string {
	values := g.MustGetAppValues(c)
	if values.UserId != "" {
		return values.UserId
	} else {
		g.log.WithFuncName("GetAppValuesOrAbort").Panic("must get user id can not get user id.")
		return ""
	}
}

func (g *GinUtils) MustGetUserId(c *gin.Context) string {
	values := g.MustGetAppValues(c)
	if values.UserId != "" {
		return values.UserId
	} else {
		g.log.WithFuncName("GetAppValuesOrAbort").Panic("must get user id can not get user id.")
		panic(errors.New("no user id"))
	}
}

func (g *GinUtils) GetCustomerId(c *gin.Context) string {
	values := g.MustGetAppValues(c)
	if values.UserId != "" {
		return values.CustomerTraceId
	} else {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("must get customer id can not get user id.")
		return ""
	}
}

func (g *GinUtils) MustGetAppContext(c *gin.Context) context2.Context {
	appContext := g.GetAppContext(c)
	if appContext == nil {
		panic(errors.New("no app context"))
	}
	return appContext
}

func (g *GinUtils) GetAppContext(c *gin.Context) context2.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		g.log.WithFuncName("GetAppValuesOrAbort").Panic("GetAppValuesOrReturn", "can not get service context for next step")
		return nil
	}
	ctx, ok := value.(context2.Context)
	if !ok {
		return nil
	}
	return ctx
}
