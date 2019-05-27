package common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
)

func (CtlCommon) uploadFile(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	fileName := c.PostForm("name")
	file, err := c.FormFile("file")
	if err != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrInternal.NewWithMsg(err.Error()))
		return
	}
	mfile, err := file.Open()
	if err != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrInternal.NewWithMsg(err.Error()))
		return
	}
	s, err := appCommon.UploadFile(ctx, ctx.UserClaims.UserId, fileName, mfile, file.Size)
	if err != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrInternal.NewWithMsg(err.Error()))
		return
	} else {
		c.JSON(200, gin.H{"path": s})
	}
}

func (CtlCommon) deleteFile(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}

}
