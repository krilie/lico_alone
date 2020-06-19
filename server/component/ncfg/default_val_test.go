package ncfg

import (
	"encoding/json"
	"testing"
)

func TestNewNConfig(t *testing.T) {
	// 默认配置
	var cfg = &Config{
		Http: Http{
			EnableSwagger: false,
			GinMode:       "debug",
			Port:          80,
			SslPri:        "",
			SslPub:        "",
			Url:           "http://localhost",
		},
		Log: Log{
			LogFile:  "log.txt",
			LogLevel: 5,
			ElfLog: &ElfLog{
				Key:    "",
				Secret: "",
				Url:    "",
			},
		},
		DB: DB{
			ConnStr:         "root:123456@tcp(localhost:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai",
			MaxOpenConn:     5,
			MaxIdleConn:     10,
			ConnMaxLeftTime: 60 * 60 * 4, // 4h
		},
		FileSave: FileSave{
			OssKey:      "",
			OssSecret:   "",
			OssEndPoint: "http://localhost/static",
			OssBucket:   "static",
			Channel:     "local",
		},
		JWT: JWT{
			NormalExpDuration: 60 * 60 * 24 * 7, // 7d
			HS256key:          "wDcD3LZl*3L$gmsDd#qSXZ2eMPcM#ps^sWWrt5*zsOoZ5hKAzrsm4&$^Tpg2PIDGoh76hEWVWkCv%cSi%aZXnyXJYC#WxWhuMBp",
		},
		Email: Email{
			Address:  "",
			Host:     "",
			Port:     465,
			UserName: "",
			Password: "",
		},
		AliSms: AliSms{
			Key:    "",
			Secret: "",
		},
	}
	cfgStr, err := json.Marshal(cfg)
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}
	println(string(cfgStr))
}
