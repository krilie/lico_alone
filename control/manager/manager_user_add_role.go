package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth_manager"
)

// /manager/user/add_role POST
//  roleId string, userId string
// role_id 角色的id
// user_id 用户的id
func ManagerUserAddRole(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	req := struct {
		RoleId string `form:"role_id" binding:"required"`
		UserId string `form:"user_id" binding:"required"`
	}{}
	err := c.ShouldBindWith(req, binding.FormPost)
	if err != nil {
		common.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(err.Error()))
		return
	}
	err = auth_manager.ManagerUserAddRole(ctx, req.UserId, req.RoleId)
	if err != nil {
		common.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
