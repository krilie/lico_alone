package ginutil

import (
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
)

// AbortWithErr abort with err use err's default http status
func (g *GinWrap) AbortWithErr(err error) {
	if nErr := errs.ToErrOrNil(err); nErr != nil {
		g.Context.AbortWithStatusJSON(200, com_model.NewRet(nErr))
	} else {
		g.Context.AbortWithStatusJSON(200, com_model.NewRetFromErr(err))
	}
}

func (g *GinWrap) AbortWithAppErr(err *errs.Err) {
	g.Context.AbortWithStatusJSON(200, com_model.NewRetFromErr(err))
}
