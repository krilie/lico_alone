package config

import (
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type config struct {
	v *viper.Viper
}

var Conf = &config{}
var log *logrus.Entry

func init() {
	log = comlog.NewLog(context.NewContext(), "br_go.common.config", "init")

	Conf.v = viper.New()

	//读取环境变量值
	Conf.v.SetEnvPrefix("LICO")
	Conf.v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	Conf.v.SetEnvKeyReplacer(replacer)

	//设置配置文件的名字
	Conf.v.SetConfigName("config")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	Conf.v.AddConfigPath("./")
	//设置配置文件类型
	Conf.v.SetConfigType("yaml")

	Conf.v.SetDefault("info.update_time", time_util.GetTimeNow().Unix()) //时间戳
	Conf.v.SetDefault("service.port", 80)
	Conf.v.SetDefault("db.ip", "localhost")
	Conf.v.SetDefault("db.port", 3306)
	Conf.v.SetDefault("db.user", "root")
	Conf.v.SetDefault("db.password", "123456")
	Conf.v.SetDefault("db.database", "br")
	Conf.v.SetDefault("db.max_open_conn", 1)
	Conf.v.SetDefault("db.max_idle_conn", 1)
	Conf.v.SetDefault("db.conn_max_left_time", 3600*7)
	Conf.v.SetDefault("ssl.public_key", "")
	Conf.v.SetDefault("ssl.private_key", "")
	Conf.v.SetDefault("jwt.normal_exp_duration", 3600*24*30)
	Conf.v.SetDefault("jwt.hs256_key", "E5Vsfs#$afasdrtfawe*^&%(")
	Conf.v.SetDefault("oss.key", "123")
	Conf.v.SetDefault("oss.secret", "123")
	Conf.v.SetDefault("oss.endpoint", "123")
	Conf.v.SetDefault("oss.bucket", "123")

	if err := Conf.v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			//err = Conf.v.WriteConfigAs("config.yaml") //new config file and ignore err
			//log.Infoln("create a new config file config.yaml at pwd path. any err:", err)
			log.Infoln("no config file use default:", err)
		default:
			log.Warnln(err)
		}
	}
}
