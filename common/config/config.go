package config

import "github.com/spf13/viper"

func init() {
	if err := LoadConfigFromYaml(Conf); err != nil {

		return
	}
}

type config struct {
	v *viper.Viper
}

var Conf = &config{}

func LoadConfigFromYaml(c *config) error {
	c.v = viper.New()
	//设置配置文件的名字
	c.v.SetConfigName("config")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	c.v.AddConfigPath("./")
	//设置配置文件类型
	c.v.SetConfigType("yaml")

	c.v.SetDefault("ok", 23)

	if err := c.v.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
