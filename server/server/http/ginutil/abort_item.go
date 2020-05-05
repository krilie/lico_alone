package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
)

// abort with err use err's default http status
func AbortWithErr(c *gin.Context, err error) {
	if nErr := errs.ToErrOrNil(err); nErr != nil {
		c.AbortWithStatusJSON(200, com_model.NewRet(nErr))
	} else {
		c.AbortWithStatusJSON(200, com_model.NewRetFromErr(err))
	}
}

func AbortWithAppErr(c *gin.Context, err *errs.Err) {
	c.AbortWithStatusJSON(200, com_model.NewRetFromErr(err))
}
