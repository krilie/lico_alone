package ctl_common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// QueryCarousel common查询轮播图
// @Summary common查询轮播图
// @Description common查询轮播图
// @Tags 公共接口
// @ID common查询轮播图
// @Produce json
// @Success 200 {object} com_model.CommonReturn{data=[]model.Carousel}
// @Failure 500 {string} errInfo
// @Router /api/common/carousel/query [GET]
func (a *CommonCtrl) QueryCarousel(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	carousel, err := a.CommonService.QueryCarousel(ctx)
	ginutil.HandlerErrorOrReturnData(c, err, carousel)
	return
}
