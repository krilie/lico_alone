package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/control/utils"
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
	panic("implement me")
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
