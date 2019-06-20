package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/comstruct"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/control/utils"
)

// /user/base/info get
// get info
func (UserCtrl) GetInfo(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	userId := ctx.GetUserId()
	if userId == "" {
		utils.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	info, err := appUser.GetInfo(ctx, userId)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, info)
		return
	}
}

// /user/base/login post
// loginName, password string
// login_name 登录名
// password 密码
func (UserCtrl) Login(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		LoginName string `form:"login_name" binding:"required,user_id"`
		Password  string `form:"password" binding:"required"`
	}{}
	if e := c.ShouldBindWith(&req, binding.FormPost); e != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	//login
	jwtString, err := appUser.Login(ctx, ctx.ClientId, req.LoginName, req.Password)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, gin.H{"token": jwtString})
		return
	}
}

// /user/base/logout post
// jwtToken string
// 从登录信息中取jwttoken
func (UserCtrl) Logout(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	logout := appUser.Logout(ctx, ctx.GetUserToken())
	if logout != nil {
		utils.ReturnWithErr(ctx, c, logout)
		return
	} else {
		c.JSON(200, comstruct.StdSuccess)
		return
	}
}

// /user/base/register POST
// loginName string, password string
// login_name 登录名
// password 密码
func (UserCtrl) Register(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		LoginName string `form:"login_name" binding:"required"`
		Password  string `form:"password" binding:"required"`
	}{}
	if e := c.ShouldBindWith(&req, binding.FormPost); e != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	//开始注册
	err := appUser.Register(ctx, req.LoginName, req.Password)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		// 可以带一次登录过程
		// jwtString, err := user_base.Login(ctx, req.LoginName, req.Password)
		c.JSON(200, comstruct.StdSuccess)
		return
	}
}

// /user/base/valid POST
// token 用户的jwttoken
// 检查这个jwtToken是否有效，并返回有效载荷
func (UserCtrl) Valid(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	token := c.PostForm("token")
	if token == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no find token param in form"))
		return
	}
	claims, e := appUser.Validate(ctx, token)
	if e != nil {
		utils.ReturnWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, claims)
		return
	}
}

// /user/base/valid_client_acc_token POST
// token
// 无权限
func (UserCtrl) ValidClientAccToken(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	token := c.PostForm("token")
	if token == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no find token param in form"))
		return
	}
	key, err := appUser.ValidateClientAccToken(ctx, token)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, key)
		return
	}
}
