package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	var param = &model.AddCatchwordModel{}
	err := ginWrap.ShouldBindBodyWith(param, binding.JSON)
	if err != nil {
		ginWrap.ReturnWithAppErr(errs.NewParamError().WithError(err).WithMsg("参数错误"))
		return
	}
	catchwordId, err := a.userService.ModuleCatchword.AddCatchword(ginWrap.AppCtx, param)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(com_model.SingleId{Id: catchwordId})
}
