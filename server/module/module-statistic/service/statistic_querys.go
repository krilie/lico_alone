package service

import (
	"context"
	"github.com/ahmetb/go-linq/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/module/module-statistic/model"
)

// QueryAllVisitorLonLat
// 查询所有访问过系统的经纬度与城市
func (a *StatisticService) QueryAllVisitorLonLat(ctx context.Context) ([]model.VisitorLonlatModel, error) {

	defer func() {
		if err := recover(); err != nil {
			a.log.Get(ctx).WithError(err).Error("panic on QueryAllVisitorLonLat")
			panic(err)
		}
	}()

	// 获取数据
	type Memo struct {
		Memo string `json:"memo" gorm:"column:memo"`
	}
	var list []Memo
	err := a.Dao.GetDb(ctx).Model(&model.StatVisitorLogs{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	a.log.Get(ctx).WithField("list_data", jsonutil.ToJson(list)).Info("get memo list")
	// 整理数据
	var listRet []model.VisitorLonlatModel
	linq.From(list).Select(func(o interface{}) interface{} {
		return o.(Memo).Memo
	}).Select(func(o interface{}) interface{} {
		return model.VisitorLonlatModel{
			Lon:  jsoniter.Get([]byte(o.(string)), "lon").ToFloat64(),
			Lat:  jsoniter.Get([]byte(o.(string)), "lat").ToFloat64(),
			City: jsoniter.Get([]byte(o.(string)), "city").ToString(),
		}
	}).Where(func(o interface{}) bool {
		return !o.(model.VisitorLonlatModel).IsEmpty()
	}).Distinct().ToSlice(&listRet)
	// 返回数据
	a.log.Get(ctx).WithField("memo_data", jsonutil.ToJson(listRet)).Info("get memo list")
	return listRet, nil
}
