package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/log"
)

// get app context or nil
func GetApplicationContextOrReturn(c *gin.Context) *context.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Error("GetApplicationContextOrReturn", "can not get application context for next step")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	contextOrNil := context.GetContextOrNil(value)
	if contextOrNil == nil {
		log.Error("GetApplicationContextOrReturn", "internal err on cast context to app context")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	return contextOrNil
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
