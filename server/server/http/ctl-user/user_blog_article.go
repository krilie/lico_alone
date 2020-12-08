package ctl_user

import (
	"github.com/gin-gonic/gin"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// ManageUpdateConfig 通过id获取文章
// @Summary 通过id获取文章
// @Description 获取一个文章信息
// @Tags 文章管理
// @ID 获取一个文章信息
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param id query string true "文章id"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/article/get_by_id [get]
func (a *UserCtrl) GetArticleById(c *gin.Context) {
	id := c.Query("id")
	article, err := a.userService.GetArticleById(a.ginUtil.MustGetAppContext(c), id)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	if article == nil {
	}

	ginutil.ReturnData(c, article)
	return
}

// ManageUpdateConfig 更新文章内容
// @Summary 更新文章内容
// @Description 更新文章内容
// @Tags 文章管理
// @ID 更新文章内容
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param article body model.UpdateArticleModel true "文章内容"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/article/update [POST]
func (a *UserCtrl) UpdateArticle(c *gin.Context) {
	log := a.log.Get(a.ginUtil.MustGetAppContext(c), "userControl", "UpdateArticle")
	param := &model.UpdateArticleModel{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		log.Errorf("user param err %v", err)
		ginutil.ReturnFailure(c, errs.ErrorParam, "参数错误")
		return
	}
	err = a.userService.UpdateArticleSample(a.ginUtil.MustGetAppContext(c), param)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}

// QueryArticle 查询文章列表
// @Summary 查询文章列表
// @Description 查询文章列表
// @Tags 文章管理
// @ID 查询文章列表
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param search_key query string true "搜索内容"
// @Param page_num query int true "页索引"
// @Param page_size query int true "页大小"
// @Success 200 {object} com_model.CommonReturn{data=com_model.PageData{data=[]model.ArticleDto}}
// @Failure 500 {string} errInfo
// @Router /api/manage/article/query [GET]
func (a *UserCtrl) QueryArticle(c *gin.Context) {
	ctx := a.ginUtil.MustGetAppContext(c)
	log := a.log.Get(ctx)
	// 参数
	var param = &struct {
		com_model.PageParams
		SearchKey string `form:"search_key" json:"search_key" xml:"search_key"`
	}{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		log.Warn(err.Error())
		ginutil.ReturnWithAppErr(c, errs.NewParamError().WithMsg(err.Error()))
		return
	}
	// 查询
	param.PageParams.CheckOkOrSetDefault()
	page, count, data, err := a.userService.ModuleArticle.QueryArticlePage(ctx, param.PageParams, param.SearchKey, a.ginUtil.MustGetUserId(c))
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, com_model.NewSuccess(com_model.PageData{
		PageInfo: com_model.PageInfo{
			TotalCount: count,
			TotalPage:  page,
			PageNum:    param.PageNum,
			PageSize:   param.PageSize,
		},
		Data: data,
	}))
}

// DeleteArticle 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags 文章管理
// @ID 删除文章
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param article_id query string true "文章id"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/article/delete [POST]
func (a *UserCtrl) DeleteArticle(c *gin.Context) {
	articleId := c.PostForm("article_id")
	if articleId == "" {
		ginutil.ReturnWithErr(c, errs.NewParamError().WithMsg("no id find on post form"))
		return
	}
	_, err := a.userService.DeleteArticleById(a.ginUtil.MustGetAppContext(c), articleId)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}

// DeleteArticle 创建文章
// @Summary 创建文章
// @Description 创建文章
// @Tags 文章管理
// @ID 创建文章
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param article body model.CreateArticleModel true "文章"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/article/create [POST]
func (a *UserCtrl) CreateArticle(c *gin.Context) {
	ctx := a.ginUtil.MustGetAppContext(c)
	log := a.log.Get(ctx)
	var param = &model.CreateArticleModel{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		log.Warn(err.Error())
		ginutil.ReturnWithAppErr(c, errs.NewParamError().WithMsg(err.Error()))
		return
	}
	err = a.userService.CreateArticle(ctx, param)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}
