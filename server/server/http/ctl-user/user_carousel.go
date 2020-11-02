package ctl_user

import (
	"github.com/gin-gonic/gin"
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
// @Success 200 {object} com_model.CommonReturn{data=[]model.Carousel}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/query [get]
func (a *UserCtrl) QueryCarousel(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	log := a.log.Get(ctx).WithFuncName("QueryCarousel")
	isOnShow := ginutil.ParseBoolOrNil(c.Query("is_on_show"))
	carousel, err := a.userService.QueryCarousel(ctx, isOnShow)
	log.Infof("err %v", err)
	ginutil.HandlerErrorOrReturnData(c, err, carousel)
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
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/create [POST]
func (a *UserCtrl) CreateCarousel(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	var item = new(model.CreateCarouselModel)
	err := c.ShouldBindJSON(item)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	err = a.userService.CreateCarousel(ctx, item)
	ginutil.HandlerErrorOrReturnSuccess(c, err)
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
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/update [POST]
func (a *UserCtrl) UpdateCarousel(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	var item = new(model.UpdateCarouselModel)
	err := c.ShouldBindJSON(item)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	err = a.userService.UpdateCarousel(ctx, item)
	ginutil.HandlerErrorOrReturnSuccess(c, err)
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
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/manage/carousel/delete_by_id [POST]
func (a *UserCtrl) DeleteCarouselById(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	carouselId := c.PostForm("carousel_id")
	err := a.userService.DeleteCarouselById(ctx, carouselId)
	ginutil.HandlerErrorOrReturnSuccess(c, err)
}
