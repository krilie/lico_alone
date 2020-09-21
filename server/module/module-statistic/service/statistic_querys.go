package service

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/module/module-statistic/model"
)

func (a *StatisticService) QueryAllVisitorLonLat(ctx context.Context) ([]model.VisitorLonlatModel, error) {
	var isInListRet = func(list []model.VisitorLonlatModel, lonlatModel model.VisitorLonlatModel) bool {
		for _, visitorLonlatModel := range list {
			if visitorLonlatModel.City == lonlatModel.City &&
				visitorLonlatModel.Lon == lonlatModel.Lon &&
				visitorLonlatModel.Lat == lonlatModel.Lat {
				return true
			}
		}
		return false
	}

	var list []struct {
		Memo string `json:"memo" gorm:"column:memo"`
	}
	err := a.Dao.GetDb(ctx).Model(&model.StatVisitorLogs{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	a.log.Get(ctx).WithField("list_data", str_util.ToJson(list)).Info("get memo list")
	var listRet []model.VisitorLonlatModel
	for i := range list {
		var item = model.VisitorLonlatModel{
			Lon:  jsoniter.Get([]byte(list[i].Memo), "lon").ToFloat64(),
			Lat:  jsoniter.Get([]byte(list[i].Memo), "lat").ToFloat64(),
			City: jsoniter.Get([]byte(list[i].Memo), "city").ToString(),
		}
		if !isInListRet(listRet, item) {
			listRet = append(listRet, item)
		}
	}
	a.log.Get(ctx).WithField("memo_data", str_util.ToJson(listRet)).Info("get memo list")
	return listRet, nil
}
