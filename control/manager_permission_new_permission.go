package control

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/control/gin_util"
	"github.com/lico603/lico_user/manager"
)

// /manager/permission/new_permission POST
// pName string, pDescription string
// name		permission的名称
// description description的描述
func ManagerPermissionNewPermission(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	name := c.PostForm("name")
	description := c.PostForm("description")
	if name == "" || description == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("name or description mast not empty"))
		return
	}
	permission, err := manager.ManagerPermissionNewPermission(ctx, name, description)
	if err != nil {
		c.JSON(200, permission)
		return
	} else {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	}
}
