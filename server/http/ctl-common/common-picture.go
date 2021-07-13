package ctl_common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// GetSlidePicById common查询单个图片信息
// @Summary common查询单个图片信息
// @Description common查询单个图片信息
// @Tags 公共接口
// @ID common查询单个图片信息
// @Produce json
// @Success 200 {object} com_model.CommonReturn{data=[]model.CarouselVo}
// @Failure 500 {string} errInfo
// @Router /api/common/picture/{id} [GET]
func (con *CommonCtrl) GetSlidePicById(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	id := ginWrap.GinCtx.Param("id")
	carousel, err := con.CommonService.ModuleCarousel.GetCarouselById(ginWrap.AppCtx, id)
	ginWrap.HandlerErrorOrReturnData(err, carousel)
	return
}
