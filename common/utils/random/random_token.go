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
func GetRandomNum(num uint) string {
	if num <= 5 {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		vcode := fmt.Sprintf("%v", rnd.Int31n(1000000))
		return vcode[:num]
	} else {
		bytes := make([]byte, 0, num)
		for i := uint(0); i < num; i++ {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			b := rnd.Intn(9) + int(byte('0'))
			bytes = append(bytes, byte(b))
		}
		return string(bytes)
	}
}
