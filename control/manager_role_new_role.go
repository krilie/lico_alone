package control

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/control/gin_util"
	"github.com/lico603/lico-my-site-user/manager"
)

// /manager/role/new_role POST
// roleName string, roleDescription string
// name	名称
// description 描述
func ManagerRoleNewRole(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrAbort(c)
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
		gin_util.AbortWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	role, e := manager.ManagerRoleNewRole(ctx, req.Name, req.Description)
	if e != nil {
		gin_util.AbortWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, role)
		return
	}
}
