package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/model/errs"
	"github.com/krilie/lico_alone/control/utils"
)

// check if request has client access token
// on header "ClientAccToken"
func CheckClientToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// take application context
		context := utils.GetAppCtxOrAbort(c)
		if context == nil {
			return
		}
		headerToken := c.GetHeader(utils.HeaderClientAccToken)
		if headerToken == "" {
			log.Info("CheckClientToken", "url no client access token", c.Request.URL)
			c.AbortWithStatusJSON(401, errs.UnAuthorized.ToStdWithMsg("no client access token find"))
			return
		}
		key, err := apiUser.ValidateClientAccToken(context, headerToken)
		if err != nil {
			utils.AbortWithErr(context, c, err)
			return
		} else {
			if key == nil {
				c.AbortWithStatusJSON(401, errs.UnAuthorized.ToStdWithMsg("client acc key not valid"))
				return
			} else {
				context.SetClientId(key.UserId)
				context.SetClientToken(key.Token)
				c.Next()
				return
			}
		}
	}
}
