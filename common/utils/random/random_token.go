package random

import (
	"encoding/base64"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

//给app用户用的random_token生成
//要求短且好看

func GetAToken() string {
	token := base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes())
	return strings.ReplaceAll(token, "=", "")
}

//获取盐值
func GetRandomNum(size int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < size; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
