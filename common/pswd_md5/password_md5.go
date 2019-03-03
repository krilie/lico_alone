package pswd_md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//对密码进行md5加密 盐值
//原密码加盐值，返回加密过后的密码,大写十六进制串
func GetMd5Password(ori string, salt string) string {
	tagPswd := md5.Sum([]byte(ori + salt + ori))
	tagPswdHex := strings.ToUpper(hex.EncodeToString(tagPswd[:]))
	return tagPswdHex
}

//检查password是否正常
func IsPasswordOk(ori, md5ed, salt string) bool {
	return GetMd5Password(ori, salt) == md5ed
}
