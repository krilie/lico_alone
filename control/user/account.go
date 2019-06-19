package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/account/pojo"
	"github.com/shopspring/decimal"
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
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		Note   string            `form:"note" binding:"required"`
		Image  string            `form:"image" binding:"-"`
		Amount decimal.Decimal   `form:"amount" binding:"required"`
		Detail []pojo.BillDetail `form:"detail" binding:"required"`
	}{}
	e := c.ShouldBindJSON(&req)
	if utils.HandlerError(ctx, c, e) {
		return
	}
	s, e := appUser.AddBill(ctx, utils.MustGetUserId(c), req.Note, req.Image, req.Amount, req.Detail)
	utils.HandlerErrorOrReturnJson(ctx, c, e, s)
}

func (UserCtrl) AddAccount(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		Name        string          `form:"name" binding:"required"`
		Num         string          `form:"num" binding:"required"`
		Description string          `form:"description" binding:"required"`
		Image       string          `form:"image" binding:"required"`
		Balance     decimal.Decimal `form:"balance" binding:"required"`
	}{}
	e := c.ShouldBindWith(&req, binding.Form)
	if utils.HandlerError(ctx, c, e) {
		return
	}
	s, e := appUser.AddAccount(ctx, utils.MustGetUserId(c), req.Name, req.Num, req.Description, req.Image, req.Balance)
	utils.HandlerErrorOrReturnJson(ctx, c, e, s)
}

func (UserCtrl) DeleteAccount(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	accountId := c.PostForm("account_id")
	if accountId == "" {
		utils.ReturnWithErr(ctx, c, errs.ErrParam.NewWithMsg("not find account id in form."))
		return
	}
	e := appUser.DeleteAccount(ctx, accountId, utils.MustGetUserId(c))
	utils.HandlerErrorOrReturnSuccess(ctx, c, e)
}

func (UserCtrl) GetMonthSummary(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	beijingTime, e := time_util.ParseBeijingTime(c.Query("time"), time_util.DefaultFormat)
	if utils.HandlerError(ctx, c, e) {
		return
	}
	summary, e := appUser.GetMonthSummary(ctx, utils.MustGetUserId(c), beijingTime)
	utils.HandlerErrorOrReturnJson(ctx, c, e, summary)
}

func (UserCtrl) GetTimeZoneSummary(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		TimeStart time.Time `form:"time_start" binding:"required" time_format:"2006-01-02 15:04:05"`
		TimeEnd   time.Time `form:"time_end" binding:"required" time_format:"2006-01-02 15:04:05"`
	}{}
	e := c.ShouldBindQuery(&req)
	if utils.HandlerError(ctx, c, e) {
		return
	}
	accountSummary, e := appUser.GetTimeZoneSummary(ctx, utils.MustGetUserId(c), req.TimeStart, req.TimeEnd)
	utils.HandlerErrorOrReturnJson(ctx, c, e, accountSummary)
}
