package random_token

import (
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"strings"
)

//给app用户用的random_token生成
//要求短且好看

func GetAToken() string {
	token := base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes())
	return strings.ReplaceAll(token, "=", "")
}
