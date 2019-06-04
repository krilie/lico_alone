package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"time"
)

// userId billId
func (UserCtrl) DeleteBill(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	billId := c.PostForm("bill_id")
	if billId == "" {
		utils.ReturnWithErr(ctx, c, errs.ErrParam.NewWithMsg("not find bill id."))
		return
	}
	err := appUser.DeleteBill(ctx, billId, utils.MustGetUserId(c))
	utils.HandlerErrorOrReturnSuccess(ctx, c, err)
}

func (UserCtrl) GetAccountHistory(c *gin.Context) {
	ctx := utils.MustGetAppCtx(c)
	req := struct {
		Start     time.Time
		End       time.Time
		AccountId string
		Note      string
	}{}
	err := c.ShouldBindWith(&req, binding.Form)
	if utils.HandlerError(ctx, c, err) {
		return
	}
	v, e := appUser.GetAccountHistory(ctx, req.Start, req.End, utils.MustGetUserId(c), req.AccountId, req.Note)
	if !utils.HandlerError(ctx, c, e) {
		c.JSON(200, v)
		return
	}
}

func (UserCtrl) GetAccountInfo(c *gin.Context) {
	panic("implement me")
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
