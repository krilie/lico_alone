package midfunc

import (
	"github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/control/gin_util"
	"github.com/lico603/lico_user/user_auth"
)

// check user has some permission request by used url
func NeedPermission(perms ...interface{}) gin.HandlerFunc {
	permsSet := mapset.NewThreadUnsafeSet()
	for e := range perms {
		permsSet.Add(e)
	}
	return func(c *gin.Context) {
		// check user get context
		ctx := gin_util.GetApplicationContextOrAbort(c)
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
		permissions, err := user_auth.UserAuthPermissions(ctx, userId)
		if err != nil {
			gin_util.AbortWithErr(ctx, c, err)
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
