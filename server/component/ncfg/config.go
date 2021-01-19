package ncfg

import (
	"github.com/krilie/lico_alone/common/run_env"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type NConfig struct {
	V      *viper.Viper
	RunEnv *run_env.RunEnv
}

func NewNConfigByCfgStr(cfgStr string) *NConfig {
	cfg := NewNConfig()
	err := cfg.LoadFromConfigTomlStr(cfgStr)
	if err != nil {
		panic(err)
	}
	return cfg
}

func NewNConfigByCfgStrFromEnv(envName string) func() *NConfig {
	return func() *NConfig {
		cfg := NewNConfig()
		cfgStr := os.Getenv(envName)
		err := cfg.LoadFromConfigTomlStr(cfgStr)
		if err != nil {
			panic(err)
		}
		return cfg
	}
}

func NewNConfigByFileFromEnv(envName string) func() *NConfig {
	return func() *NConfig {
		cfg := NewNConfig()
		filePath := os.Getenv(envName)
		err := cfg.LoadConfigByFile(filePath)
		if err != nil {
			panic(err)
		}
		return cfg
	}
}

func NewNConfig() *NConfig {
	var cfg = &NConfig{V: viper.New(), RunEnv: run_env.RunEnvLocal}

	//读取环境变量值
	cfg.V.SetEnvPrefix("MYAPP")
	cfg.V.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	cfg.V.SetEnvKeyReplacer(replacer)

	err := cfg.LoadDefaultConfig()
	if err != nil {
		log.Println("加载默认配置值时失败 :" + err.Error())
	}

	return cfg
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
