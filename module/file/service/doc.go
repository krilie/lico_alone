package service

import (
	"github.com/krilie/lico_alone/common/config"
)

var ossKey = config.GetString("oss.key")
var ossSecret = config.GetString("oss.secret")
var ossEndPoint = config.GetString("oss.endpoint")
var ossBucket = config.GetString("oss.bucket")
