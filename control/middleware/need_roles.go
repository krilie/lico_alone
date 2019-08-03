package middleware

import (
	"github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/model/errs"
	"github.com/krilie/lico_alone/control/utils"
)

func NeedRoles(roles ...interface{}) gin.HandlerFunc {
	roleSet := mapset.NewThreadUnsafeSet()
	for e := range roles {
		roleSet.Add(e)
	}
	return func(c *gin.Context) {
		// check user get context
		ctx := utils.GetAppCtxOrAbort(c)
		if ctx == nil {
			return
		}
		//get user id from context
		userId := ctx.GetUserId()
		if userId == "" {
			c.AbortWithStatusJSON(errs.UnAuthorized.HttpStatus, errs.UnAuthorized.ToStdWithMsg("can not get user id from app context,check login status"))
			return
		}
		//check user has permission
		userRoles, err := apiAuthUser.GetRoles(ctx, userId)
		if err != nil {
			utils.AbortWithErr(ctx, c, err)
			return
		}
		sect := roleSet.Intersect(userRoles)
		if sect.Cardinality() == 0 {
			utils.AbortWithAppErr(ctx, c, errs.ErrNoPermission.NewWithMsg("no permission"))
			return
		}
		c.Next()
	}
}
