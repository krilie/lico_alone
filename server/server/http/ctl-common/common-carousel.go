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
// @Success 200 {object} com_model.CommonReturn{data=[]model.CarouselDto}
// @Failure 500 {string} errInfo
// @Router /api/common/carousel/query [GET]
func (con *CommonCtrl) QueryCarousel(c *gin.Context) {
	ctx := con.ginUtil.MustGetAppValues(c)
	carousel, err := con.CommonService.QueryCarousel(ctx)
	ginutil.HandlerErrorOrReturnData(c, err, carousel)
	return
}
