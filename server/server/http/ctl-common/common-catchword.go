package ctl_common

//// QueryCatchword common查询时代热词
//// @Summary common查询时代热词
//// @Description common查询时代热词
//// @Tags 公共接口
//// @ID common查询时代热词
//// @Produce json
//// @Success 200 {object} com_model.CommonReturn{data=[]model.CatchwordVo}
//// @Failure 500 {string} errInfo
//// @Router /api/common/catchword/query [GET]
//func (con *CommonCtrl) QueryCatchword(c *gin.Context) {
//	ctx := con.ginUtil.MustGetAppContext(c)
//
//	con.CommonService.ModuleCatchword.Dao.QueryListForWebShow(ctx,,c.GetQuery("key_word"),c.Query("form"),c.GetInt("limit"))
//
//	ginutil.HandlerErrorOrReturnData(c, err, carousel)
//	return
//}
