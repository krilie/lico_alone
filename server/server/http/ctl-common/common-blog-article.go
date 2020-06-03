package ctl_common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// QueryArticleSample 查询文章列表
// @Summary 查询文章列表
// @Description 查询文章列表
// @Tags 公共接口
// @ID 查询文章列表
// @Produce json
// @Param search_key query string true "搜索内容"
// @Param page_num query int true "页索引"
// @Param page_size query int true "页大小"
// @Success 200 {object} com_model.CommonReturn{Data=com_model.PageData{Data=[]model.QueryArticleModel}}
// @Failure 500 {string} errInfo
// @Router /api/common/article/query_sample [GET]
func (a *CommonCtrl) QueryArticleSample(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	log := a.log.Get(ctx)
	var param = &struct {
		com_model.PageParams
		SearchKey string `json:"search_key"`
	}{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		log.Warn(err.Error())
		ginutil.ReturnWithAppErr(c, errs.NewParamError().WithMsg(err.Error()))
		return
	}
	pageData, err := a.CommonService.QueryArticleSamplePage(ctx, param.PageParams, param.SearchKey)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, com_model.NewSuccess(pageData))
	return
}
