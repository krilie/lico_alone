package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/string_util"
	"github.com/krilie/lico_alone/control/utils"
	"time"
)

// /manager/client_user/create/new_acc_token POST
// 要有管理员权限 admin client权限 [如果是client权限要求登录者和目标用户是同一用户]
// target_user_id 目标用户id,给哪个用户生成acc_token,
// description string 这个key的描述
// exp time.Time 这个key的过期时间,utx时间戳
// 逻辑层做了参数检查，不用在这里检查参数，一些非空检查是必要的
func (ManagerCtrl) CreateNewAccToken(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
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
	} else if num, err := string_util.GetInt64(exp); err != nil {
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
