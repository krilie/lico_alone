package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth_manager"
)

// /manager/role/new_permission Post
// roleId string, permissionId string
// role_id role的id
// permission_id permission的id
func ManagerRoleAddPermission(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	roleId := c.PostForm("role_id")
	permissionID := c.PostForm("permission_id")
	if roleId == "" || permissionID == "" {
		common.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("role_id or permission_id must exists"))
	}
	err := auth_manager.ManagerRoleAddPermission(ctx, roleId, permissionID)
	if err != nil {
		common.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
