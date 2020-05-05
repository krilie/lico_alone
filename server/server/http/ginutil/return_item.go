package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
)

// 处理错误，如果有错误返回真 无错误返回假
func HandlerError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	} else {
		ReturnWithErr(c, err)
		return true
	}
}

// 处理错误 如果没有返回通用成功
func HandlerErrorOrReturnSuccess(c *gin.Context, err error) {
	if err == nil {
		c.JSON(200, com_model.StdSuccess)
		return
	} else {
		ReturnWithErr(c, err)
		return
	}
}

// 处理错误 如果没有返回通用成功
func HandlerErrorOrReturnJson(c *gin.Context, err error, ret interface{}) {
	if err == nil {
		c.JSON(200, ret)
		return
	} else {
		ReturnWithErr(c, err)
		return
	}
}

// abort with err use err's default http status
func ReturnWithErr(c *gin.Context, err error) {
	if nErr := errs.ToErrOrNil(err); nErr != nil {
		c.JSON(200, com_model.NewRet(nErr))
	} else {
		c.JSON(200, com_model.NewRetFromErr(err))
	}
}

func ReturnWithAppErr(c *gin.Context, err *errs.Err) {
	c.JSON(200, com_model.NewRet(err))
}

func ReturnOk(c *gin.Context) {
	c.JSON(200, com_model.StdSuccess)
}
func ReturnFailure(code errs.ErrCode, c *gin.Context) {
	c.JSON(200, com_model.NewFailure(code, ""))
}
