package middleware

import (
	"github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

// check user has some permission request by used url
func NeedPermission(perms ...interface{}) gin.HandlerFunc {
	permsSet := mapset.NewThreadUnsafeSet()
	for e := range perms {
		permsSet.Add(e)
	}
	return func(c *gin.Context) {
		// check user get context
		ctx := common.GetApplicationContextOrAbort(c)
		if ctx == nil {
			return
		}
		//get user id from context
		userId := ctx.GetUserIdOrEmpty()
		if userId == "" {
			c.AbortWithStatusJSON(errs.UnAuthorized.HttpStatus, errs.UnAuthorized.ToStdWithMsg("can not get user id from app context,check login status"))
			return
		}
		//check user has permission
		permissions, err := auth.UserAuthPermissions(ctx, userId)
		if err != nil {
			common.AbortWithErr(ctx, c, err)
			return
		}
		sect := permsSet.Intersect(permissions)
		if sect.Cardinality() == 0 {
			c.AbortWithStatusJSON(errs.ErrNoPermission.HttpStatus, errs.ErrNoPermission.ToStdWithMsg("no permission"))
			return
		}
		c.Next()
	}
}
