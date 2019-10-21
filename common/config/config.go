package config

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"strings"
)

var (
	v   *viper.Viper
	Cfg Config
)

func init() {
	log := clog.NewLog(context.Background(), "lico_alone.common.config", "init")

	//var ConfigFilePath string
	//flag.StringVar(&ConfigFilePath, "config", "./config.yml", "config file path")
	//flag.Parse()

	v = viper.New()

	//读取环境变量值
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	// 共同配置
	v.SetDefault("gin_mode", "debug") //时间戳
	v.SetDefault("http_port", 80)
	v.SetDefault("enable_swagger", true)
	v.SetDefault("ssl_pri", "")
	v.SetDefault("ssl_pub", "")
	// db 配置
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.port", 3306)
	v.SetDefault("db.user", "root")
	v.SetDefault("db.password", "123456")
	v.SetDefault("db.db_name", "lico")
	v.SetDefault("db.max_open_conn", 1)
	v.SetDefault("db.max_idle_conn", 1)
	v.SetDefault("db.conn_max_left_time", 3600*7)
	// jwt 配置
	v.SetDefault("jwt.normal_exp_duration", 3600*24*30)
	v.SetDefault("jwt.hs256_key", "E5Vsfs#$afasdrtfawe*^&%(")
	v.SetDefault("jwt.private_key_path", "E5Vsfs#$afasdrtfawe*^&%(")
	// 文件保存配置
	v.SetDefault("file_save.local_file_save_dir", "static_files")
	v.SetDefault("file_save.local_file_save_url", "http://localhost/static_files")
	v.SetDefault("file_save.save_type", "local")
	v.SetDefault("file_save.oss_key", "local")
	v.SetDefault("file_save.oss_secret", "local")
	v.SetDefault("file_save.oss_end_point", "local")
	v.SetDefault("file_save.oss_bucket", "local")
	// email 配置
	v.SetDefault("email.host", "aaa")
	v.SetDefault("email.port", 32)
	v.SetDefault("email.user_name", "aaa")
	v.SetDefault("email.password", "aaa")
	v.SetDefault("email.address", "aaa")

	if err := LoadConfigByFile("./config.yaml"); err != nil {
		log.Error(err.Error())
		return
	}
}

func LoadConfigByFile(name string) error {
	v.SetConfigFile(name)
	if err := v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			//err = Conf.v.WriteConfigAs("config.yaml") //new config file and ignore err
			log.Infoln("no config file use default:", err)
		default:
			log.Warnln(err)
		}
	}

	err := v.Unmarshal(&Cfg)
	if err != nil {
		return err
	}
	return nil
}

/**

db:
  conn_max_left_time: 25200
  db_name: lico
  host: rpi.lizo.top
  max_idle_conn: 1
  max_open_conn: 1
  password: "123456"
  port: 3306
  user: root
email:
  address: aaa
  host: aaa
  password: aaa
  port: 32
  user_name: aaa
enable_swagger: true
file_save:
  local_file_save_dir: static_files
  local_file_save_url: http://localhost/static_files
  oss_bucket: local
  oss_end_point: local
  oss_key: local
  oss_secret: local
  save_type: local
gin_mode: debug
http_port: 80
jwt:
  hs256_key: E5Vsfs#$afasdrtfawe*^&%(
  normal_exp_duration: 2592000
  private_key_path: E5Vsfs#$afasdrtfawe*^&%(
ssl_pri: ""
ssl_pub: ""


*/
