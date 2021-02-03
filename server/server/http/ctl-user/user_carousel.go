package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/utils/strutil"
	"github.com/krilie/lico_alone/module/module-carousel/model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// QueryCarousel 管理者查询轮播图
// @Summary 管理者查询轮播图
// @Description 管理者查询轮播图
// @Tags 轮播图
// @ID 管理者查询轮播图
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param is_on_show query bool false "是否显示"
// @Success 200 {object} com_model.CommonReturn{data=[]model.CarouselDto}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/query [get]
func (a *UserCtrl) QueryCarousel(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	log := a.log.Get(ginWrap.AppCtx).WithFuncName("QueryCarousel")
	isOnShow := strutil.ParseBoolOrNil(c.Query("is_on_show"))
	carousel, err := a.userService.QueryCarousel(ginWrap.AppCtx, isOnShow)
	log.Infof("err %v", err)
	ginWrap.HandlerErrorOrReturnData(err, carousel)
	return
}

// CreateCarousel 管理员创建轮播图
// @Summary 管理员创建轮播图
// @Description 管理员创建轮播图
// @Tags 轮播图
// @ID 管理员创建轮播图
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param carousel body model.CreateCarouselModel true "单个文件"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/create [POST]
func (a *UserCtrl) CreateCarousel(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var item = new(model.CreateCarouselModel)
	err := c.ShouldBindJSON(item)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	err = a.userService.CreateCarousel(ginWrap.AppCtx, item)
	ginWrap.HandlerErrorOrReturnSuccess(err)
	return
}

// UpdateCarousel 管理员更新轮播图
// @Summary 管理员更新轮播图
// @Description 管理员更新轮播图
// @Tags 轮播图
// @ID 管理员更新轮播图
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param carousel body model.UpdateCarouselModel true "更新结构"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/update [POST]
func (a *UserCtrl) UpdateCarousel(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var item = new(model.UpdateCarouselModel)
	err := c.ShouldBindJSON(item)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	err = a.userService.UpdateCarousel(ginWrap.AppCtx, item)
	ginWrap.HandlerErrorOrReturnSuccess(err)
	return
}

// DeleteCarouselById 管理员删除轮播图
// @Summary 管理员删除轮播图
// @Description 管理员删除轮播图
// @Tags 轮播图
// @ID 管理员删除轮播图
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param carousel_id formData string true "id"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/delete_by_id [POST]
func (a *UserCtrl) DeleteCarouselById(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	carouselId := c.PostForm("carousel_id")
	err := a.userService.DeleteCarouselById(ginWrap.AppCtx, carouselId)
	ginWrap.HandlerErrorOrReturnSuccess(err)
}
