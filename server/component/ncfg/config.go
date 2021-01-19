package ncfg

import (
	"flag"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type NConfig struct {
	V *viper.Viper
}

func NewNConfigByCfgStr(cfgStr string) *NConfig {
	cfg := NewNConfig()
	err := cfg.LoadFromConfigTomlStr(cfgStr)
	if err != nil {
		panic(err)
	}
	return cfg
}

func NewNConfigByCfgStrFromEnv(envName string) *NConfig {
	cfg := NewNConfig()
	cfgStr := os.Getenv(envName) //"MYAPP_TEST_CONFIG"
	err := cfg.LoadFromConfigTomlStr(cfgStr)
	if err != nil {
		panic(err)
	}
	return cfg
}

func NewNConfigByCfgStrFromEnvTest() *NConfig {
	return NewNConfigByCfgStrFromEnv("MYAPP_TEST_CONFIG")
}

func NewNConfig() *NConfig {
	var cfg = &NConfig{V: viper.New()}

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
	cfg.TryLoadFromCmdConfigFile()

	return cfg
}

func (cfg *NConfig) TryLoadFromCmdConfigFile() bool {
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
	open, err := os.Open(name)
	if err != nil {
		return err
	}
	defer open.Close()
	cfgStr, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	return cfg.LoadFromConfigTomlStr(string(cfgStr))
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
		return nil
	}
}

func (cfg *NConfig) LoadDefaultConfig() error {
	err := cfg.LoadFromConfigTomlStr(defaultCfg)
	if err != nil {
		return err
	}
	return nil
}
