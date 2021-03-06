package ctl_user

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-file/model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

type UpdateFileReturn struct {
	Url    string `json:"url" swag:"true,file's url'"`
	Bucket string `json:"bucket" swag:"true,file's bucket'"`
	Key    string `json:"key" swag:"true,file's key'"`
}

// 文件上传
// 文件删除
// 文件查询

// UploadFile 文件上传
// @Summary 文件上传
// @Description 文件上传
// @Tags 文件管理
// @ID 文件上传
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param file formData file true "单个文件"
// @Success 200 {object} com_model.CommonReturn{data=UpdateFileReturn}
// @Failure 500 {string} errInfo
// @Router /api/manage/file/upload [POST]
func (a *UserCtrl) UploadFile(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	//NOTE: 使用临时文件缓存
	err := c.Request.ParseMultipartForm(1024 * 1024 * 5) // 5mb放在内存中 超过放在临时文件中
	if err != nil {
		ginWrap.ReturnFailure(errs.ErrorInternal, err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		ginWrap.ReturnFailure(errs.ErrorParam, "no file found")
		return
	}
	f, err := file.Open()
	if err != nil {
		ginWrap.ReturnFailure(errs.ErrorInternal, err.Error())
		return
	}
	defer f.Close()
	values := context.MustGetAppValues(ginWrap.AppCtx)
	url, bucket, key, err := a.userService.UploadFile(ginWrap.AppCtx, values.UserId, file.Filename, f, int(file.Size))
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(&UpdateFileReturn{Url: url, Bucket: bucket, Key: key})
	return

	//NOTE: 不使用文件缓存
	//reader, err := c.Request.MultipartReader()
	//if err != nil {
	//	ginWrap.ReturnWithErr( err)
	//	return
	//}
	//for {
	//	p, err := reader.NextPart()
	//	if err == io.EOF {
	//		ginWrap.ReturnWithErr( errs.NewParamError().WithMsg("no file"))
	//		return
	//	}
	//	if err != nil {
	//		ginWrap.ReturnWithErr( err)
	//		return
	//	}
	//
	//	name := p.FormName()
	//	if name == "" {
	//		continue
	//	}
	//
	//	if name != "file" {
	//		continue
	//	}
	//
	//	filename := p.FileName()
	//	if filename == "" {
	//		continue
	//	}
	//
	//	size := str_util.GetIntOrDef(p.Header.Get("size"), -1)
	//
	//	url, bucket, key, err := a.userService.UploadFile(ginWrap.AppCtx, ctx.UserId, p.FileName(), p, size)
	//	if err != nil {
	//		ginWrap.ReturnWithErr(err)
	//		return
	//	}
	//	ginWrap.ReturnData(&UpdateFileReturn{Url: url, Bucket: bucket, Key: key})
	//	return
	//}

}

// DeleteFile 文件删除
// @Summary 文件删除
// @Description 文件删除
// @Tags 文件管理
// @ID 文件删除
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param file_id formData string true "文件记录id"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/file/delete [POST]
func (a *UserCtrl) DeleteFile(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	log := a.log.Get(ginWrap.AppCtx).WithFuncName("DeleteFile")
	fileId := c.PostForm("file_id")
	if fileId == "" {
		log.Errorf("no file id found file_id %v", fileId)
		ginWrap.ReturnFailure(errs.ErrorParam, "no file id found")
		return
	}
	err := a.userService.ModuleFile.DeleteFileById(ginWrap.AppCtx, fileId)
	ginWrap.HandlerErrorOrReturnSuccess(err)
	return
}

// QueryFile 文件查询
// @Summary 文件查询
// @Description 文件查询
// @Tags 文件管理
// @ID 文件查询
// @Produce json
// @Param Authorization header string true "凭证token" default({{token}})
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
// @Success 200 {object} com_model.CommonReturn{data=com_model.PageData{data=[]model.FileMasterDto}}
// @Failure 500 {string} errInfo
// @Failure 500 {object} com_model.CommonReturn{data=object}
// @Router /api/manage/file/query [GET]
func (a *UserCtrl) QueryFile(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var param = &model.QueryFileParam{}
	err := c.BindQuery(param)
	if err != nil {
		ginWrap.ReturnFailure(errs.ErrorParam, err.Error())
		return
	}
	totalPage, totalCount, pageNum, pageSize, files, err := a.userService.ModuleFile.QueryFilePage(ginWrap.AppCtx, *param)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(com_model.PageData{
		PageInfo: com_model.PageInfo{
			TotalCount: totalCount,
			TotalPage:  totalPage,
			PageNum:    pageNum,
			PageSize:   pageSize,
		},
		Data: linq.From(files).Select(func(o interface{}) interface{} {
			return o.(*model.FileMaster).ToDto()
		}).Results(),
	})
	return
}
