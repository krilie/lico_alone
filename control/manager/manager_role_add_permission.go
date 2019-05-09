package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module/user_auth/user_auth_manager"
)

// /manager/role/new_permission Post
// roleId string, permissionId string
// role_id role的id
// permission_id permission的id
func ManagerRoleAddPermission(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	roleId := c.PostForm("role_id")
	permissionID := c.PostForm("permission_id")
	if roleId == "" || permissionID == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("role_id or permission_id must exists"))
	}
	err := user_auth_manager.ManagerRoleAddPermission(ctx, roleId, permissionID)
	if err != nil {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}