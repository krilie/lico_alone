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
	V      *viper.Viper
	Cfg    *Config
	RunEnv *run_env.RunEnv
}

func NewNConfig() *NConfig {
	var cfg = &NConfig{V: viper.New(), Cfg: &Config{}, RunEnv: run_env.RunEnvLocal}

	//读取环境变量值
	cfg.V.SetEnvPrefix("MYAPP")
	cfg.V.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	cfg.V.SetEnvKeyReplacer(replacer)

	err := cfg.LoadDefaultConfig()
	if err != nil {
		log.Println("加载默认配置值时失败 :" + err.Error())
	}

	// 默认位置
	if err := cfg.LoadConfigByFile("config.toml"); err != nil {
		log.Println("未能从默认位置加载配置 :" + err.Error())
	}

	// 有没有环境变量配置配置文件
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
	cfg.V.SetConfigFile(name)
	if err := cfg.V.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			err = cfg.V.WriteConfigAs("config.toml") //new config file and ignore err
			log.Println("no config file gen and use default:", err)
		default:
			log.Println(err)
		}
		return err
	} else {
		err := cfg.V.Unmarshal(cfg.Cfg)
		if err != nil {
			return err
		}
		return nil
	}
}

func (cfg *NConfig) LoadFromConfigTomlStr(cfgStr string) error {
	cfg.V.SetConfigType("toml")
	if err := cfg.V.MergeConfig(strings.NewReader(cfgStr)); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("no config find on cfg str gen and use default:", err)
		default:
			log.Println(err)
		}
		return err
	} else {
		err := cfg.V.Unmarshal(cfg.Cfg)
		if err != nil {
			return err
		}
		return nil
	}
}

func (cfg *NConfig) LoadDefaultConfig() error {
	err2 := cfg.LoadFromConfigTomlStr(defaultCfg)
	if err2 != nil {
		return err2
	}
	err := cfg.V.Unmarshal(cfg.Cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *NConfig) GetInt(key string) int {
	ok := cfg.V.IsSet(key)
	if !ok {
		return 0
	}
	return cfg.V.GetInt(key)
}

func (cfg *NConfig) GetString(key string) string {
	ok := cfg.V.IsSet(key)
	if !ok {
		return ""
	}
	return cfg.V.GetString(key)
}
