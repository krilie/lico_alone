package model

type VisitorLonlatModel struct {
	Lon  float64 `json:"lon"`  // 经度
	Lat  float64 `json:"lat"`  // 纬度
	City string  `json:"city"` // 城市
}

func (v VisitorLonlatModel) IsEmpty() bool {
	return v.Lat == 0 && v.Lon == 0 && v.City == ""
}
