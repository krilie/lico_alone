package config

import "log"
import "github.com/spf13/viper"


type config struct {
	v  *viper.Viper;
}


func LoadConfigFromYaml (c *config) error  {
	c.v = viper.New();

	//设置配置文件的名字
	c.v.SetConfigName("config")

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	c.v.AddConfigPath("%GOPATH/src/")
	c.v.AddConfigPath("./")

	//设置配置文件类型
	c.v.SetConfigType("yaml");

	if err := c.v.ReadInConfig(); err != nil{
		return  err;
	}

	log.Printf("age: %s, name: %s \n", c.v.Get("information.age"), c.v.Get("information.name"));
	return nil;
}



