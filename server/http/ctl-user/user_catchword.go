package ctl_user

import (
	"github.com/gin-gonic/gin"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-catchword/model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// AddCatchword 添加时代语
// @Tags 时代语
// @ID 添加时代语
// @Summary 添加时代语
// @Description 添加时代语
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param param body model.AddCatchwordModel true "添加内容"
// @Success 200 {object} com_model.CommonReturn{data=com_model.SingleId}
// @Failure 500 {string} errInfo
// @Router /api/manage/catchword/add [POST]
func (a *UserCtrl) AddCatchword(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var param model.AddCatchwordModel
	err := ginWrap.ShouldBindJSON(&param)
	if err != nil {
		ginWrap.ReturnWithAppErr(errs.NewParamError().WithError(err).WithMsg("参数错误"))
		return
	}

	catchwordId, err := a.userService.ModuleCatchword.AddCatchword(ginWrap.AppCtx, &param)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}

	ginWrap.ReturnData(com_model.SingleId{Id: catchwordId})
}

// UpdateCatchword 更新时代语
// @Tags 时代语
// @ID 更新时代语
// @Summary 更新时代语
// @Description 更新时代语
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param param body model.UpdateCatchwordModel true "添加内容"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {string} errInfo
// @Router /api/manage/catchword/update [POST]
func (a *UserCtrl) UpdateCatchword(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var param model.UpdateCatchwordModel
	err := ginWrap.ShouldBindJSON(&param)
	if err != nil {
		ginWrap.ReturnWithAppErr(errs.NewParamError().WithError(err).WithMsg("参数错误"))
		return
	}

	err = a.userService.ModuleCatchword.UpdateCatchword(ginWrap.AppCtx, &param)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}

	ginWrap.ReturnOk()
}

// UpdateCatchword 删除时代语
// @Tags 时代语
// @ID 删除时代语
// @Summary 删除时代语
// @Description 删除时代语
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param id formData string true "id"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {string} errInfo
// @Router /api/manage/catchword/delete [POST]
func (a *UserCtrl) DeleteCatchword(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	err := a.userService.ModuleCatchword.DeleteCatchword(ginWrap.AppCtx, ginWrap.PostForm("id"))
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}

	ginWrap.ReturnOk()
}

// UpdateCatchword 查询时代语
// @Tags 时代语
// @ID 查询时代语
// @Summary 查询时代语
// @Description 查询时代语
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param key_word query string true "key word"
// @Param page_size query int true "page size"
// @Param page_num query int true "page num"
// @Success 200 {object} com_model.CommonReturn{data=com_model.PageData{data=[]model.CatchwordVo}}
// @Failure 500 {string} errInfo
// @Router /api/manage/catchword/query [GET]
func (a *UserCtrl) QueryCatchword(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var param = struct {
		KeyWord string `json:"key_word" url:"key_word" form:"key_word" binding:""`
		com_model.PageParams
	}{}
	err := ginWrap.ShouldBindQuery(&param)
	if err != nil {
		ginWrap.ReturnWithParamErr(err)
		return
	}

	pageInfo, data, err := a.userService.ModuleCatchword.Dao.QueryList(ginWrap.AppCtx, param.KeyWord, com_model.PageParams{
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	})
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}

	ginWrap.ReturnData(com_model.NewPageData2(pageInfo, data))
}
