package config

import (
	"github.com/lico603/lico_user/common/log"
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

	Conf.v.SetDefault("service.port", 8080)
	Conf.v.SetDefault("db.ip", "192.168.31.238")
	Conf.v.SetDefault("db.port", 3306)
	Conf.v.SetDefault("db.user", "root")
	Conf.v.SetDefault("db.password", "123456")
	Conf.v.SetDefault("db.database", "user")
	Conf.v.SetDefault("db.max_open_conn", 1)
	Conf.v.SetDefault("db.max_idle_conn", 1)
	Conf.v.SetDefault("db.conn_max_left_time", 3600*7)
	Conf.v.SetDefault("jwt.normal_exp_duration", 3600*24*30)
	//ssl key
	//Conf.v.SetDefault("ssl.public_key", "public.key")
	Conf.v.SetDefault("ssl.public_key", "")
	//Conf.v.SetDefault("ssl.private_key", "private.key")
	Conf.v.SetDefault("ssl.private_key", "")
	Conf.v.SetDefault("jwt.hs256_key", "E5Vsfs#$afasdrtfawe*^&%(")

	if err := Conf.v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			//err = Conf.v.WriteConfigAs("config.yaml") //new config file and ignore err
			//log.Log.Infoln("create a new config file config.yaml at pwd path. any err:", err)
			log.Infoln("no config file use default:", err)
		default:
			log.Log.Warnln(err)
		}
	}
}
