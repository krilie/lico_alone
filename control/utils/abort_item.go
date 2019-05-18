package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
)

// get app context or nil
func GetApplicationContextOrAbort(c *gin.Context) *context_util.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Error("GetApplicationContextOrAbort", "can not get application context for next step")
		c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	contextOrNil := context_util.GetContextOrNil(value)
	if contextOrNil == nil {
		log.Error("GetApplicationContextOrAbort", "internal err on cast context to app context")
		c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	return contextOrNil
}

// abort with err use err's default http status
func AbortWithErr(ctx *context_util.Context, c *gin.Context, err error) {
	if errLocal, ok := err.(*errs.Error); ok {
		c.AbortWithStatusJSON(errLocal.HttpStatus, errLocal.ToStdReturn())
	} else {
		c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdAppendMsg(err.Error()))
	}
}

func AbortWithAppErr(ctx *context_util.Context, c *gin.Context, err *errs.Error) {
	c.AbortWithStatusJSON(err.HttpStatus, err.ToStdReturn())
}
