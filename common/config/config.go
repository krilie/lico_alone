package config

import (
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/spf13/viper"
)

type config struct {
	v *viper.Viper
}

var Conf = &config{}

func init() {
	Conf.v = viper.New()
	//设置配置文件的名字
	Conf.v.SetConfigName("config")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	Conf.v.AddConfigPath("./")
	//设置配置文件类型
	Conf.v.SetConfigType("yaml")

	Conf.v.SetDefault("ok", 23)
	Conf.v.SetDefault("ok2", 24)
	Conf.v.SetDefault("service.port", 80)

	if err := Conf.v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			err = Conf.v.WriteConfigAs("config.yaml") //new config file and ignore err
			log.Log.Infoln("create a new config file config.yaml at pwd path. any err:", err)
		default:
			log.Log.Warnln(err)
		}
	}
}
