package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/control/utils"
	"time"
)

// userId billId
func (UserCtrl) DeleteBill(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	billId := c.PostForm("bill_id")
	if !validator.IsIdStr(billId) {
		utils.ReturnWithErr(ctx, c, errs.ErrParam.NewWithMsg("not find bill id."))
		return
	}
	err := appUser.DeleteBill(ctx, billId, utils.MustGetUserId(c))
	utils.HandlerErrorOrReturnSuccess(ctx, c, err)
}

func (UserCtrl) GetAccountHistory(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		Start     time.Time `form:"start" binding:"required" time_format:"2006-01-02 15:04:05"`
		End       time.Time `form:"end" binding:"required" time_format:"2006-01-02 15:04:05"`
		AccountId string    `form:"account_id" binding:"required,id_str"`
		Note      string    `form:"note" binding:"-"`
	}{}
	err := c.ShouldBindWith(&req, binding.Form)
	if utils.HandlerError(ctx, c, err) {
		return
	}
	v, e := appUser.GetAccountHistory(ctx, req.Start, req.End, utils.MustGetUserId(c), req.AccountId, req.Note)
	utils.HandlerErrorOrReturnJson(ctx, c, e, v)
}

func (UserCtrl) GetAccountInfo(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	userId := c.Query("user_id")
	if !validator.IsIdStr(userId) {
		utils.ReturnWithErr(ctx, c, errs.ErrParam.NewWithMsg("not find user id"))
		return
	}
	infos, e := appUser.GetAccountInfo(ctx, userId)
	utils.HandlerErrorOrReturnJson(ctx, c, e, infos)
}

func (UserCtrl) AddBill(c *gin.Context) {
	panic("implement me")
}

func (UserCtrl) AddAccount(c *gin.Context) {
	panic("implement me")
}

func (UserCtrl) DeleteAccount(c *gin.Context) {
	panic("implement me")
}

func (UserCtrl) GetMonthSummary(c *gin.Context) {
	panic("implement me")
}

func (UserCtrl) GetTimeZoneSummary(c *gin.Context) {
	panic("implement me")
}
