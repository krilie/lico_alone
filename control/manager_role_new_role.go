package control

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module_user_auth/user_auth_manager"
)

// /manager/role/new_role POST
// roleName string, roleDescription string
// name	名称
// description 描述
func ManagerRoleNewRole(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	//匿名结构体，参数
	req := &struct {
		Name        string `binding:"required" form:"name"`
		Description string `binding:"required" form:"description"`
	}{}
	e := c.ShouldBindWith(req, binding.FormPost)
	if e != nil {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	role, e := user_auth_manager.ManagerRoleNewRole(ctx, req.Name, req.Description)
	if e != nil {
		gin_util.ReturnWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, role)
		return
	}
}
