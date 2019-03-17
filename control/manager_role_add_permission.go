package control

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/common_struct"
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/control/gin_util"
	"github.com/lico603/lico-my-site-user/manager"
)

// /manager/role/new_permission Post
// roleId string, permissionId string
// role_id role的id
// permission_id permission的id
func ManagerRoleAddPermission(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrAbort(c)
	if ctx == nil {
		return
	}
	roleId := c.PostForm("role_id")
	permissionID := c.PostForm("permission_id")
	if roleId == "" || permissionID == "" {
		gin_util.AbortWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("role_id or permission_id must exists"))
	}
	err := manager.ManagerRoleAddPermission(ctx, roleId, permissionID)
	if err != nil {
		gin_util.AbortWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
