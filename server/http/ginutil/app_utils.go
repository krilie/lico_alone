package ginutil

import (
	context2 "context"
	"errors"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
)

// GetAppValuesOrAbort 给中间件使用
// get app context or nil
func (g *GinWrap) GetAppValuesOrAbort() *context.AppCtxValues {
	appContext := g.GetAppContext()
	if appContext == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("can not get service context for next step")
		g.Context.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	value := context.GetAppValues(appContext)
	if value == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("can not get service context for next step")
		g.Context.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	return value
}

func (g *GinWrap) GetAppContextOrAbort() context2.Context {
	appContext := g.GetAppContext()
	if appContext == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("can not get service context for next step")
		g.Context.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInternal()))
		return nil
	}
	return appContext
}

func (g *GinWrap) GetUserIdOrAbort() string {
	values := g.GetAppValuesOrAbort()
	if values == nil {
		return ""
	}
	if values.UserId == "" {
		g.AbortWithErr(errs.NewInvalidToken().WithMsg("登录信息无效"))
		return ""
	} else {
		return values.UserId
	}
}

// GetAppValuesOrReturn get app context or nil
func (g *GinWrap) GetAppValuesOrReturn() *context.AppCtxValues {
	appContext := g.GetAppContext()
	if appContext == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("GetAppValuesOrReturn", "can not get service context for next step")
		g.Context.JSON(200, com_model.NewRet(errs.NewInternal().WithMsg("ctx not get")))
		return nil
	}
	value := context.GetAppValues(appContext)
	if value == nil {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("GetAppValuesOrReturn", "can not get service context for next step")
		g.Context.JSON(200, com_model.NewRet(errs.NewInternal().WithMsg("ctx not get")))
		return nil
	}
	return value
}

func (g *GinWrap) MustGetAppValues() *context.AppCtxValues {
	appContext := g.MustGetAppContext()
	value := context.MustGetAppValues(appContext)
	return value
}

func (g *GinWrap) GetUserId() string {
	values := g.MustGetAppValues()
	if values.UserId != "" {
		return values.UserId
	} else {
		g.log.WithFuncName("GetAppValuesOrAbort").Panic("must get user id can not get user id.")
		return ""
	}
}

func (g *GinWrap) MustGetUserId() string {
	values := g.MustGetAppValues()
	if values.UserId != "" {
		return values.UserId
	} else {
		g.log.WithFuncName("GetAppValuesOrAbort").Panic("must get user id can not get user id.")
		panic(errors.New("no user id"))
	}
}

func (g *GinWrap) GetCustomerId() string {
	values := g.MustGetAppValues()
	if values.CustomerTraceId != "" {
		return values.CustomerTraceId
	} else {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("must get customer id can not get user id.")
		return ""
	}
}

func (g *GinWrap) MustGetAppContext() context2.Context {
	appContext := g.GetAppContext()
	if appContext == nil {
		panic(errors.New("no app context"))
	}
	return appContext
}

func (g *GinWrap) GetAppContext() context2.Context {
	value, exists := g.Context.Get(GinKeyAppContext)
	if !exists {
		g.log.WithFuncName("GetAppValuesOrAbort").Error("GetAppValuesOrReturn", "can not get service context for next step")
		return nil
	}
	ctx, ok := value.(context2.Context)
	if !ok {
		return nil
	}
	return ctx
}
