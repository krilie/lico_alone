package gin_util

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/common/context_util"
	"github.com/lico603/lico_user/common/log"
)

// get app context or nil
func GetApplicationContextOrReturn(c *gin.Context) *context_util.Context {
	value, exists := c.Get(GinKeyAppContext)
	if !exists {
		log.Error("GetApplicationContextOrReturn", "can not get application context for next step")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	contextOrNil := context_util.GetContextOrNil(value)
	if contextOrNil == nil {
		log.Error("GetApplicationContextOrReturn", "internal err on cast context to app context")
		c.JSON(500, errs.ErrInternal.ToStdReturn())
		return nil
	}
	return contextOrNil
}

// abort with err use err's default http status
func ReturnWithErr(ctx *context_util.Context, c *gin.Context, err error) {
	if errLocal, ok := err.(*errs.Error); ok {
		c.JSON(errLocal.HttpStatus, errLocal.ToStdReturn())
	} else {
		c.JSON(500, errs.ErrInternal.ToStdAppendMsg(err.Error()))
	}
}

func ReturnWithAppErr(ctx *context_util.Context, c *gin.Context, err *errs.Error) {
	c.JSON(err.HttpStatus, err.ToStdReturn())
}
