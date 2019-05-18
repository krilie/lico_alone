package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth_manager"
)

// /manager/role/new_role POST
// roleName string, roleDescription string
// name	名称
// description 描述
func ManagerRoleNewRole(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
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
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	role, e := apiManagerUser.NewRole(ctx, req.Name, req.Description)
	if e != nil {
		utils.ReturnWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, role)
		return
	}
}
