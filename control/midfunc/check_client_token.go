package midfunc

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/string_util"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module/user_auth/user_base"
)

// check if request has client access token
// on header "ClientAccToken"
func CheckClientToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// take application context
		context := gin_util.GetApplicationContextOrAbort(c)
		if context == nil {
			return
		}
		headerToken := c.GetHeader(gin_util.HeaderClientAccToken)
		if headerToken == "" {
			log.Info("CheckClientToken", "url no client access token", c.Request.URL)
			c.AbortWithStatusJSON(401, errs.UnAuthorized.ToStdWithMsg("no client access token find"))
			return
		}
		key, err := user_base.UserValidateClientAccToken(context, headerToken)
		if err != nil {
			gin_util.AbortWithErr(context, c, err)
			return
		} else {
			context.ClientId = string_util.NewString(key.UserId)
			context.ClientAccToken = string_util.NewString(key.Token)
			c.Next()
			return
		}
	}
}
