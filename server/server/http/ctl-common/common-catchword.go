package ctl_common

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// QueryCatchword common查询时代热词
// @Summary common查询时代热词
// @Description common查询时代热词
// @Tags 公共接口
// @ID common查询时代热词
// @Produce json
// @Param key_word query string true "搜索关键词"
// @Param form query string true "从什么时间开始 2021-02-06 10:34:03"
// @Param limit query int true "倒序取多少个"
// @Success 200 {object} com_model.CommonReturn{data=[]model.CatchwordVo} "时间排序"
// @Failure 500 {string} errInfo
// @Router /api/common/catchword/query [GET]
func (con *CommonCtrl) QueryCatchword(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, con.log)

	var data, err = con.CommonService.
		ModuleCatchword.
		Dao.
		QueryListForWebShow(
			ginWrap.AppCtx,
			ginWrap.QueryParamOrEmpty("key_word"),
			now.MustParse(ginWrap.QueryParamOrEmpty("form")),
			c.GetInt("limit"))
	ginWrap.HandlerErrorOrReturnData(err, data)
	return
}
