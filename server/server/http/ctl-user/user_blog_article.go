package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// ManageUpdateConfig 通过id获取文章
// @Summary 通过id获取文章
// @Description 获取一个文章信息
// @Tags 文章
// @ID 获取一个文章信息
// @Produce  json
// @Param id query string true "文章id"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {string} errInfo
// @Router /api/manage/article/get_by_id [get]
func (a *UserCtrl) GetArticleById(c *gin.Context) {
	id := c.Query("id")
	article, err := a.userService.QueryArticleById(ginutil.MustGetAppCtx(c), id)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, article)
	return
}

// ManageUpdateConfig 通过id获取文章
// @Summary 通过id获取文章
// @Description 获取一个文章信息
// @Tags 文章
// @ID 获取一个文章信息
// @Produce  json
// @Param id query string true "文章id"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {string} errInfo
// @Router /api/manage/article/update [POST]
func (a *UserCtrl) UpdateArticle(c *gin.Context) {
	log := a.log.NewWithCtx(ginutil.MustGetAppCtx(c), "userControl", "UpdateArticle")
	param := &model.UpdateArticleModel{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		log.Errorf("user param err %v", err)
		ginutil.ReturnFailure(c, errs.ErrorParam, "参数错误")
		return
	}
	err = a.userService.UpdateArticle(ginutil.MustGetAppCtx(c), id)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}
