package config

import (
	"fmt"
	"testing"
)

func TestConfig(T *testing.T) {
	i := Conf.v.GetInt("service.port")
	fmt.Println(i)
	fmt.Println(Conf.v.GetInt("ok"))
	fmt.Println(Conf.v.GetInt("ok2"))
}
