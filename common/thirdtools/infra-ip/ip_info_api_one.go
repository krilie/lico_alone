package infra_ip

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
{
	"status": "success",
	"country": "美国",
	"countryCode": "US",
	"region": "CA",
	"regionName": "加利福尼亚州",
	"city": "洛杉矶",
	"zip": "90009",
	"lat": 34.0522,
	"lon": -118.244,
	"timezone": "America/Los_Angeles",
	"isp": "Psychz Networks",
	"org": "Psychz Networks",
	"as": "AS40676 Psychz Networks",
	"query": "45.34.23.12"
}
*/
type IpInfo struct {
	As          string  `json:"as"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Message     string  `json:"message"`
	Isp         string  `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

type IpInfoApiOne struct {
	Url string
}

func (i *IpInfoApiOne) GetIpInfoRegionCityOrEmpty(ctx context.Context, ip string) string {
	info, err := i.GetIpInfo(ctx, ip)
	if err != nil {
		return ""
	}
	return info.RegionName + "-" + info.City
}

func NewIpInfoApiOne() *IpInfoApiOne {
	return &IpInfoApiOne{Url: "http://ip-api.com/json"}
}

func (i *IpInfoApiOne) GetIpInfo(ctx context.Context, ip string) (*IpInfoCommon, error) {
	req, err := http.Get(fmt.Sprintf("%v/%v?lang=zh-CN", i.Url, ip))
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(req.Body)
	if req.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	var ipInfo = &IpInfo{}
	err = json.Unmarshal(body, ipInfo)
	if err != nil {
		return nil, err
	}
	if ipInfo.Status == "success" {
		return &IpInfoCommon{
			RegionName:  ipInfo.RegionName,
			City:        ipInfo.City,
			Country:     ipInfo.Country,
			Lat:         ipInfo.Lat,
			Lon:         ipInfo.Lon,
			RawResponse: string(body),
		}, nil
	} else {
		return nil, errors.New(ipInfo.Message)
	}
}
