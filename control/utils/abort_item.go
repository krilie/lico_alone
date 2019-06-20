package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	lcontext "github.com/krilie/lico_alone/common/context"
)

// 给中间件使用
// get app context or nil
func GetAppCtxOrAbort(c *gin.Context) *lcontext.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Error("GetAppCtxOrAbort", "can not get application context for next step")
		c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	contextOrNil := lcontext.GetContextOrNil(value)
	if contextOrNil == nil {
		log.Error("GetAppCtxOrAbort", "internal err on cast context to app context")
		c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	return contextOrNil
}

// abort with err use err's default http status
func AbortWithErr(ctx context.Context, c *gin.Context, err error) {
	if errLocal, ok := err.(*errs.Error); ok {
		c.AbortWithStatusJSON(errLocal.HttpStatus, errLocal.ToStdReturn())
	} else {
		c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdAppendMsg(err.Error()))
	}
}

func AbortWithAppErr(ctx context.Context, c *gin.Context, err *errs.Error) {
	c.AbortWithStatusJSON(err.HttpStatus, err.ToStdReturn())
}
