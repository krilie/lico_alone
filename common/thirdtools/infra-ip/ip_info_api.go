package infra_ip

import "context"

type IpInfoCommon struct {
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	RawResponse string  `json:"raw_response"`
}

type IIpInfo interface {
	GetIpInfo(ctx context.Context, ip string) (*IpInfoCommon, error)
	GetIpInfoRegionCityOrEmpty(ctx context.Context, ip string) string
}
