package uuid_util

import (
	"encoding/hex"
	"github.com/satori/go.uuid"
)

//没有-的uuid 用做主键
//小写十六进制串 32个
func GetUuid() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
