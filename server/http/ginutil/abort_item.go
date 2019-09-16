package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/errs"
)

// abort with err use err's default http status
func AbortWithErr(c *gin.Context, err error) {
	if lerr := errs.ToErrOrNil(err); lerr != nil {
		c.AbortWithStatusJSON(lerr.Code, cmodel.NewRet(lerr))
	} else {
		c.AbortWithStatusJSON(500, cmodel.RetFromErr(err))
	}
}

func AbortWithAppErr(c *gin.Context, err *errs.Err) {
	c.AbortWithStatusJSON(err.Code, cmodel.RetFromErr(err))
}
