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
// @ID 查询文章列表简单
// @Produce json
// @Param search_key query string true "搜索内容"
// @Param page_num query int true "页索引"
// @Param page_size query int true "页大小"
// @Success 200 {object} com_model.CommonReturn{data=com_model.PageData{data=[]model.QueryArticleModel}}
// @Failure 500 {string} errInfo
// @Router /api/common/article/query_sample [GET]
func (con *CommonCtrl) QueryArticleSample(c *gin.Context) {
	ctx := con.ginUtil.MustGetAppCtx(c)
	log := con.log.Get(ctx)
	var param = &struct {
		com_model.PageParams
		SearchKey string `form:"search_key" json:"search_key" xml:"search_key"  binding:"required"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		log.Warn(err.Error())
		ginutil.ReturnWithAppErr(c, errs.NewParamError().WithMsg(err.Error()))
		return
	}
	param.PageParams.CheckOkOrSetDefault()
	pageData, err := con.CommonService.QueryArticleSamplePage(ctx, param.PageParams, param.SearchKey)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, pageData)
	return
}

// GetArticle 获取article
// @Summary 获取article
// @Description 获取article
// @Tags 公共接口
// @ID 获取article
// @Produce json
// @Param article_id query string true "搜索内容"
// @Success 200 {object} com_model.CommonReturn{data=model.ArticleDto}
// @Failure 500 {string} errInfo
// @Router /api/common/article/get_article [GET]
func (con *CommonCtrl) GetArticle(c *gin.Context) {
	ctx := con.ginUtil.MustGetAppCtx(c)
	log := con.log.Get(ctx)
	articleId := c.Query("article_id")

	article, err := con.CommonService.GetArticleById(ctx, articleId)
	if err != nil {
		log.Error(err)
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, article)
	return
}
