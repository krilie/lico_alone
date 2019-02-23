package config

import "github.com/lico603/lico-my-site-user/common/log"

func GetInt(key string) int {
	ok := Conf.v.IsSet(key)
	if !ok {
		log.Fatal("key not find in config file or set.")
		return 0
	}
	return Conf.v.GetInt(key)
}
func GetString(key string) string {
	ok := Conf.v.IsSet(key)
	if !ok {
		log.Fatal("key not find in config file or set.")
		return ""
	}
	return Conf.v.GetString(key)
}
