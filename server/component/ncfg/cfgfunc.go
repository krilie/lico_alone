package ncfg

import (
	"flag"
	"github.com/krilie/lico_alone/common/run_env"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type NConfig struct {
	v      *viper.Viper
	Cfg    *Config
	RunEnv *run_env.RunEnv
}

func NewNConfig() *NConfig {
	var cfg = &NConfig{v: viper.New(), Cfg: &Config{}, RunEnv: run_env.RunEnvLocal}

	//读取环境变量值
	cfg.v.SetEnvPrefix("MYAPP")
	cfg.v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	cfg.v.SetEnvKeyReplacer(replacer)

	err := cfg.LoadDefaultConfig()
	if err != nil {
		log.Println("加载默认配置值时失败 :" + err.Error())
	}

	// 默认位置
	if err := cfg.LoadConfigByFile("config.toml"); err != nil {
		log.Println("未能从默认位置加载配置 :" + err.Error())
	}

	// 没有环境变量
	configFile := os.Getenv("APP_CONFIG_PATH")
	if configFile != "" {
		// 加载配置文件
		if err := cfg.LoadConfigByFile(configFile); err != nil {
			log.Println("未能从环境变量加载配置文件 :" + err.Error())
		} else {
			log.Println("已从环境变量加载配置文件 :" + configFile)
		}
	} else {
		log.Println("未能找到环境变量中的配置路径 不从环境变量加载配置")
	}

	// 命令行配置文件优先级最高
	cfg.TryLoadFromArgConfigFile()

	return cfg
}

func (cfg *NConfig) TryLoadFromArgConfigFile() bool {
	// The default set of command-line flags, parsed from os.Args.
	file, _ := os.Open("/dev/null")
	var commandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	commandLine.SetOutput(file)
	var configFilePath = commandLine.String("config", "", "配置文件")
	err := commandLine.Parse(os.Args[1:])
	if err != nil {
		log.Printf("命令行参数没有config %v", err)
	}
	if *configFilePath != "" {
		_ = cfg.LoadConfigByFile(*configFilePath)
		return true
	}
	return false
}

func (cfg *NConfig) LoadConfigByFile(name string) error {
	cfg.v.SetConfigFile(name)
	if err := cfg.v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			err = cfg.v.WriteConfigAs("config.toml") //new config file and ignore err
			log.Println("no config file gen and use default:", err)
		default:
			log.Println(err)
		}
		return err
	} else {
		err := cfg.v.Unmarshal(cfg.Cfg)
		if err != nil {
			return err
		}
		return nil
	}
}

func (cfg *NConfig) LoadDefaultConfig() error {
	cfg.v.SetDefault("http.enable_swagger", false)
	cfg.v.SetDefault("http.gin_mode", "debug")
	cfg.v.SetDefault("http.port", 80)
	cfg.v.SetDefault("http.ssl_pri", "")
	cfg.v.SetDefault("http.ssl_pub", "")
	cfg.v.SetDefault("http.url", "http://localhost")
	cfg.v.SetDefault("log.log_file", "log.txt")
	cfg.v.SetDefault("log.log_level", 5)
	cfg.v.SetDefault("log.elf_log.key", "")
	cfg.v.SetDefault("log.elf_log.secret", "")
	cfg.v.SetDefault("log.elf_log.url", "")
	cfg.v.SetDefault("db.conn_str", "root:123456@tcp(localhost:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	cfg.v.SetDefault("db.max_open_conn", 5)
	cfg.v.SetDefault("db.max_idle_conn", 10)
	cfg.v.SetDefault("db.conn_max_left_time", 14400)
	cfg.v.SetDefault("file_save.oss_key", "")
	cfg.v.SetDefault("file_save.oss_secret", "")
	cfg.v.SetDefault("file_save.oss_end_point", "http://localhost/static")
	cfg.v.SetDefault("file_save.oss_bucket", "static")
	cfg.v.SetDefault("file_save.channel", "local")
	cfg.v.SetDefault("jwt.normal_exp_duration", 604800)
	cfg.v.SetDefault("jwt.hs_256_key", "wDcD3LZl*3L$gmsDd#qSXZ2eMPcM#ps^sWWrt5*zsOoZ5hKAzrsm4&$^Tpg2PIDGoh76hEWVWkCv%cSi%aZXnyXJYC#WxWhuMBp")
	cfg.v.SetDefault("email.address", "")
	cfg.v.SetDefault("email.host", "")
	cfg.v.SetDefault("email.port", 465)
	cfg.v.SetDefault("email.user_name", "")
	cfg.v.SetDefault("email.password", "")
	cfg.v.SetDefault("ali_sms.key", "")
	cfg.v.SetDefault("ali_sms.secret", "")

	err := cfg.v.Unmarshal(cfg.Cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *NConfig) GetInt(key string) int {
	ok := cfg.v.IsSet(key)
	if !ok {
		return 0
	}
	return cfg.v.GetInt(key)
}

func (cfg *NConfig) GetString(key string) string {
	ok := cfg.v.IsSet(key)
	if !ok {
		return ""
	}
	return cfg.v.GetString(key)
}
