package ctl_user

import (
	"github.com/gin-gonic/gin"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-file/model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// 文件上传
// 文件删除
// 文件查询

// UpdateFile 文件上传
// @Summary 文件上传
// @Description 文件上传
// @Tags 文件管理
// @ID 文件上传
// @Produce json
// @Param file formData file true "单个文件"
// @Success 200 {object} com_model.CommonReturn{data=UpdateFileReturn}
// @Failure 500 {string} errInfo
// @Router /api/manage/file/upload [POST]
func (a *UserCtrl) UpdateFile(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	file, err := c.FormFile("file")
	if err != nil {
		ginutil.ReturnFailure(c, errs.ErrorParam, "no file found")
		return
	}
	f, err := file.Open()
	if err != nil {
		ginutil.ReturnFailure(c, errs.ErrorInternal, err.Error())
		return
	}
	defer f.Close()
	url, bucket, key, err := a.userService.UploadFile(ctx, ctx.UserId, file.Filename, f, int(file.Size))
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, &UpdateFileReturn{Url: url, Bucket: bucket, Key: key})
	return
}

type UpdateFileReturn struct {
	Url    string `json:"url" swag:"true,file's url'"`
	Bucket string `json:"bucket" swag:"true,file's bucket'"`
	Key    string `json:"key" swag:"true,file's key'"`
}

// DeleteFile 文件删除
// @Summary 文件删除
// @Description 文件删除
// @Tags 文件管理
// @ID 文件删除
// @Produce json
// @Param file_id formData string true "文件记录id"
// @Success 200 {object} com_model.CommonReturn{}
// @Failure 500 {string} errInfo
// @Router /api/manage/file/delete [POST]
func (a *UserCtrl) DeleteFile(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	fileId := c.PostForm("file_id")
	if fileId == "" {
		ginutil.ReturnFailure(c, errs.ErrorParam, "no file id found")
		return
	}
	err := a.userService.ModuleFile.DeleteFileById(ctx, fileId)
	ginutil.HandlerErrorOrReturnSuccess(c, err)
	return
}

// QueryArticle 文件查询
// @Summary 文件查询
// @Description 文件查询
// @Tags 文件管理
// @ID 文件查询
// @Produce json
// @Param page_num query int true "page_num页索引"
// @Param page_size query int true "page_size页大小"
// @Param key_name_like formData string true "key_name_like"
// @Param bucket_name_like formData string true "bucket_name_like"
// @Param url_like formData string true "url_like"
// @Param user_id formData string true "user_id"
// @Param biz_type formData string true "biz_type"
// @Param content_type formData string true "content_type"
// @Param created_at_begin formData string true "created_at_begin"
// @Param created_at_end formData string true "created_at_end"
// @Success 200 {object} com_model.CommonReturn{com_model.PageData{data=[]model.FileMaster}}
// @Failure 500 {string} errInfo
// @Router /api/manage/file/query [POST]
func (a *UserCtrl) QueryFile(c *gin.Context) {
	ctx := ginutil.MustGetAppCtx(c)
	var param = &model.QueryFileParam{}
	err := c.BindQuery(param)
	if err != nil {
		ginutil.ReturnFailure(c, errs.ErrorParam, err.Error())
		return
	}
	totalPage, totalCount, pageNum, pageSize, files, err := a.userService.ModuleFile.QueryFilePage(ctx, *param)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, com_model.PageData{
		PageInfo: com_model.PageInfo{
			TotalCount: totalCount,
			TotalPage:  totalPage,
			PageNum:    pageNum,
			PageSize:   pageSize,
		},
		Data: files,
	})
	return
}
