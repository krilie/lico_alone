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
	ginWrap := ginutil.NewGinWrap(c, con.log)

	log := con.log.Get(ginWrap.AppCtx)
	log.Info("on query article sample")
	var param = &struct {
		com_model.PageParams
		SearchKey string `form:"search_key" json:"search_key" xml:"search_key"  binding:"required"`
	}{}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		log.Warn(err.Error())
		ginWrap.ReturnWithAppErr(errs.NewParamError().WithMsg(err.Error()))
		return
	}
	param.PageParams.CheckOkOrSetDefault()
	pageData, err := con.CommonService.QueryArticleSamplePage(ginWrap.AppCtx, param.PageParams, param.SearchKey, ginWrap.GetCustomerId())
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(pageData)
	return
}

// GetArticle 获取article
// @Summary 获取article
// @Description 获取article
// @Tags 公共接口
// @ID 获取article
// @Produce json
// @Param article_id query string true "文章id"
// @Success 200 {object} com_model.CommonReturn{data=model.ArticleDto}
// @Failure 500 {string} errInfo
// @Router /api/common/article/get_article [GET]
func (con *CommonCtrl) GetArticle(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	log := con.log.Get(ginWrap.AppCtx)
	articleId := c.Query("article_id")

	article, err := con.CommonService.GetArticleById(ginWrap.AppCtx, articleId, ginWrap.GetCustomerId())
	if err != nil {
		log.Error(err)
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(article)
	return
}

// MarkArticleLike 文章点like
// @Summary 文章点like
// @Description 文章点like
// @Tags 公共接口
// @ID 文章点like
// @Produce json
// @Param article_id formData string true "article id"
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/common/article/mark/like [POST]
func (con *CommonCtrl) MarkArticleLike(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	log := con.log.Get(ginWrap.AppCtx)
	articleId := c.PostForm("article_id")
	customerId := ginWrap.MustGetAppValues().CustomerTraceId

	err := con.CommonService.ModuleArticle.AddLike(ginWrap.AppCtx, customerId, articleId)
	if err != nil {
		log.Error(err)
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnOk()
	return
}

// MarkArticleDisLike 文章点dislike
// @Summary 文章点dislike
// @Description 文章点dislike
// @Tags 公共接口
// @ID 文章点dislike
// @Produce json
// @Param article_id formData string true "article id"
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/common/article/mark/dislike [POST]
func (con *CommonCtrl) MarkArticleDisLike(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	log := con.log.Get(ginWrap.AppCtx)
	articleId := c.PostForm("article_id")
	customerId := ginWrap.MustGetAppValues().CustomerTraceId

	err := con.CommonService.ModuleArticle.AddDisLike(ginWrap.AppCtx, customerId, articleId)
	if err != nil {
		log.Error(err)
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnOk()
	return
}

// RemoveMarkArticleLike 文章点like-remove
// @Summary 文章点like-remove
// @Description 文章点like-remove
// @Tags 公共接口
// @ID 文章点like-remove
// @Produce json
// @Param article_id body string true "id"
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/common/article/mark/remove_like [POST]
func (con *CommonCtrl) RemoveMarkArticleLike(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	log := con.log.Get(ginWrap.AppCtx)
	articleId := c.PostForm("article_id")
	customerId := ginWrap.MustGetAppValues().CustomerTraceId

	err := con.CommonService.ModuleArticle.RemoveLike(ginWrap.AppCtx, customerId, articleId)
	if err != nil {
		log.Error(err)
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnOk()
	return
}

// RemoveMarkArticleDisLike 文章点dislike-remove
// @Summary 文章点dislike-remove
// @Description 文章点dislike-remove
// @Tags 公共接口
// @ID 文章点dislike-remove
// @Produce json
// @Param article_id query string true "article id"
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/common/article/mark/remove_dislike [POST]
func (con *CommonCtrl) RemoveMarkArticleDisLike(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	log := con.log.Get(ginWrap.AppCtx)
	articleId := c.PostForm("article_id")
	customerId := ginWrap.MustGetAppValues().CustomerTraceId

	err := con.CommonService.ModuleArticle.RemoveDisLike(ginWrap.AppCtx, customerId, articleId)
	if err != nil {
		log.Error(err)
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnOk()
	return
}
