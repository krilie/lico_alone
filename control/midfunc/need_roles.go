package midfunc

import (
	"github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module/user_auth/user_auth"
)

func NeedRoles(roles ...interface{}) gin.HandlerFunc {
	roleSet := mapset.NewThreadUnsafeSet()
	for e := range roles {
		roleSet.Add(e)
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
		userRoles, err := user_auth.UserAuthRoles(ctx, userId)
		if err != nil {
			gin_util.AbortWithErr(ctx, c, err)
			return
		}
		sect := roleSet.Intersect(userRoles)
		if sect.Cardinality() == 0 {
			gin_util.AbortWithAppErr(ctx, c, errs.ErrNoPermission.NewWithMsg("no permission"))
			return
		}
		c.Next()
	}
}
