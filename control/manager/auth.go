package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/comstruct"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/control/utils"
	"time"
)

// /manager/role/new_permission Post
// roleId string, permissionId string
// role_id role的id
// permission_id permission的id
func (ManagerCtrl) AddPermissionToRole(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	roleId := c.PostForm("role_id")
	permissionID := c.PostForm("permission_id")
	if roleId == "" || permissionID == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("role_id or permission_id must exists"))
	}
	err := appManage.AddPermissionToRole(ctx, roleId, permissionID)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, comstruct.StdSuccess)
		return
	}
}

// /manager/user/add_role POST
//  roleId string, userId string
// role_id 角色的id
// user_id 用户的id
func (ManagerCtrl) AddRoleToUser(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		RoleId string `form:"role_id" binding:"required"`
		UserId string `form:"user_id" binding:"required"`
	}{}
	err := c.ShouldBindWith(&req, binding.FormPost)
	if err != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(err.Error()))
		return
	}
	err = appManage.AddRoleToUser(ctx, req.UserId, req.RoleId)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, comstruct.StdSuccess)
		return
	}
}

// /manager/client_user/create/new_acc_token POST
// 要有管理员权限 admin client权限 [如果是client权限要求登录者和目标用户是同一用户]
// target_user_id 目标用户id,给哪个用户生成acc_token,
// description string 这个key的描述
// exp time.Time 这个key的过期时间,utx时间戳
// 逻辑层做了参数检查，不用在这里检查参数，一些非空检查是必要的
func (ManagerCtrl) CreateNewAccToken(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	//取参数
	targetUserId := c.PostForm("target_user_id")
	if targetUserId == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("target_user_id must exists"))
		return
	}
	description := c.PostForm("description")
	if description == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("description must exists"))
		return
	}
	var expNum int64
	exp := c.PostForm("exp")
	if exp == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("exp must exists"))
		return
	} else if num, err := str_util.GetInt64(exp); err != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("exp must a utx num"))
		return
	} else {
		expNum = num //默认中国上海时区
	}
	//添加新的token
	token, err := appManage.NewClientAccToken(ctx, targetUserId, description, time.Unix(expNum, 0))
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		//200
		c.JSON(200, token)
		return
	}
}

// /manager/permission/new_permission POST
// pName string, pDescription string
// name		permission的名称
// description description的描述
func (ManagerCtrl) CreateNewPermission(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	name := c.PostForm("name")
	description := c.PostForm("description")
	if name == "" || description == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("name or description mast not empty"))
		return
	}
	permission, err := appManage.NewPermission(ctx, name, description)
	if err != nil {
		c.JSON(200, permission)
		return
	} else {
		utils.ReturnWithErr(ctx, c, err)
		return
	}
}

// /manager/role/new_role POST
// roleName string, roleDescription string
// name	名称
// description 描述
func (ManagerCtrl) CreateNewRole(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	//匿名结构体，参数
	req := &struct {
		Name        string `binding:"required" form:"name"`
		Description string `binding:"required" form:"description"`
	}{}
	e := c.ShouldBindWith(&req, binding.FormPost)
	if e != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	role, e := appManage.NewRole(ctx, req.Name, req.Description)
	if e != nil {
		utils.ReturnWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, role)
		return
	}
}
