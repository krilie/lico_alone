package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/clog"
	lcontext "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/model"
	"github.com/krilie/lico_alone/common/model/errs"
)

// get app context or nil
func GetAppCtxOrReturn(c *gin.Context) *lcontext.Context {
	log = log.WithField(clog.Function, "GetAppCtxOrReturn")
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Error("GetAppCtxOrReturn", "can not get application context for next step")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	contextOrNil := lcontext.GetContextOrNil(value)
	if contextOrNil == nil {
		log.Error("GetAppCtxOrReturn", "internal err on cast context to app context")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	return contextOrNil
}

func MustGetAppCtx(c *gin.Context) *lcontext.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Panic("GetAppCtxOrReturn", "can not get application context for next step")
		return nil
	}
	return lcontext.MustGetContext(value)
}

func MustGetUserId(c *gin.Context) string {
	ctx := MustGetAppCtx(c)
	if ctx.GetUserId() != "" {
		return ctx.GetUserId()
	} else {
		log.Panic("must get user id can not get user id.")
		return ""
	}
}

// 处理错误，如果有错误返回真 无错误返回假
func HandlerError(ctx context.Context, c *gin.Context, err error) bool {
	if err == nil {
		return false
	} else {
		ReturnWithErr(ctx, c, err)
		return true
	}
}

// 处理错误 如果没有返回通用成功
func HandlerErrorOrReturnSuccess(ctx context.Context, c *gin.Context, err error) {
	if err == nil {
		c.JSON(200, model.StdSuccess)
		return
	} else {
		ReturnWithErr(ctx, c, err)
		return
	}
}

// 处理错误 如果没有返回通用成功
func HandlerErrorOrReturnJson(ctx context.Context, c *gin.Context, err error, ret interface{}) {
	if err == nil {
		c.JSON(200, ret)
		return
	} else {
		ReturnWithErr(ctx, c, err)
		return
	}
}

// abort with err use err's default http status
func ReturnWithErr(ctx context.Context, c *gin.Context, err error) {
	if errLocal, ok := err.(*errs.Error); ok {
		c.JSON(errLocal.HttpStatus, errLocal.ToStdReturn())
	} else {
		c.JSON(500, errs.ErrInternal.ToStdAppendMsg(err.Error()))
	}
}

func ReturnWithAppErr(ctx context.Context, c *gin.Context, err *errs.Error) {
	c.JSON(err.HttpStatus, err.ToStdReturn())
}
