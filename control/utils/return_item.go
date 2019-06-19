package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
)

// get app context or nil
func GetAppCtxOrReturn(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Error("GetAppCtxOrReturn", "can not get application context for next step")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	contextOrNil := context.GetContextOrNil(value)
	if contextOrNil == nil {
		log.Error("GetAppCtxOrReturn", "internal err on cast context to app context")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	return contextOrNil
}

func MustGetAppCtx(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Panic("GetAppCtxOrReturn", "can not get application context for next step")
		return nil
	}
	return context.MustGetContext(value)
}

func MustGetUserId(c *gin.Context) string {
	ctx := MustGetAppCtx(c)
	if ctx.UserClaims != nil && ctx.UserClaims.UserId != "" {
		return ctx.UserClaims.UserId
	} else {
		log.Panic("must get user id can not get user id.")
		return ""
	}
}

// 处理错误，如果有错误返回真 无错误返回假
func HandlerError(ctx *context.Context, c *gin.Context, err error) bool {
	if err == nil {
		return false
	} else {
		ReturnWithErr(ctx, c, err)
		return true
	}
}

// 处理错误 如果没有返回通用成功
func HandlerErrorOrReturnSuccess(ctx *context.Context, c *gin.Context, err error) {
	if err == nil {
		c.JSON(200, comstruct.StdSuccess)
		return
	} else {
		ReturnWithErr(ctx, c, err)
		return
	}
}

// 处理错误 如果没有返回通用成功
func HandlerErrorOrReturnJson(ctx *context.Context, c *gin.Context, err error, ret interface{}) {
	if err == nil {
		c.JSON(200, ret)
		return
	} else {
		ReturnWithErr(ctx, c, err)
		return
	}
}

// abort with err use err's default http status
func ReturnWithErr(ctx *context.Context, c *gin.Context, err error) {
	if errLocal, ok := err.(*errs.Error); ok {
		c.JSON(errLocal.HttpStatus, errLocal.ToStdReturn())
	} else {
		c.JSON(500, errs.ErrInternal.ToStdAppendMsg(err.Error()))
	}
}

func ReturnWithAppErr(ctx *context.Context, c *gin.Context, err *errs.Error) {
	c.JSON(err.HttpStatus, err.ToStdReturn())
}
