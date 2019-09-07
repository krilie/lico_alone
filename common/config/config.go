package config

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var (
	v   *viper.Viper
	Cfg Config
)

func init() {
	log := clog.NewLog(context.Background(), "lico_alone.common.config", "init")

	v = viper.New()

	//读取环境变量值
	v.SetEnvPrefix("LICO")
	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	//设置配置文件的名字
	v.SetConfigName("config")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	v.AddConfigPath("./")
	//设置配置文件类型
	v.SetConfigType("yml")

	v.SetDefault("info.update_time", time.Now().Unix()) //时间戳
	v.SetDefault("service.port", 80)
	v.SetDefault("db.host", "127.0.0.1")
	v.SetDefault("db.port", 3306)
	v.SetDefault("db.user", "root")
	v.SetDefault("db.password", "123456")
	v.SetDefault("db.database", "lico")
	v.SetDefault("db.max_open_conn", 1)
	v.SetDefault("db.max_idle_conn", 1)
	v.SetDefault("db.conn_max_left_time", 3600*7)
	v.SetDefault("ssl.public_key", "")
	v.SetDefault("ssl.private_key", "")
	v.SetDefault("jwt.normal_exp_duration", 3600*24*30)
	v.SetDefault("jwt.hs256_key", "E5Vsfs#$afasdrtfawe*^&%(")
	v.SetDefault("oss_s3.key", "123")
	v.SetDefault("oss_s3.secret", "123")
	v.SetDefault("oss_s3.endpoint", "123")
	v.SetDefault("oss_s3.bucket", "123")
	v.SetDefault("file_save.local_file_save_dir", "static_files")
	v.SetDefault("file_save.local_file_save_url", "http://localhost/static_files")
	v.SetDefault("file_save.save_type", "local")

	if err := v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			//err = Conf.v.WriteConfigAs("config.yaml") //new config file and ignore err
			//log.Infoln("create a new config file config.yaml at pwd path. any err:", err)
			log.Infoln("no config file use default:", err)
		default:
			log.Warnln(err)
		}
	}

	err := v.Unmarshal(&Cfg)
	if err != nil {
		log.Panic(err)
	}

}
