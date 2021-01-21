package id_util

import (
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomStrWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomStr(length int) string {
	return RandomStrWithCharset(length, charset)
}

func RandomInt() int64 {
	return rand.Int63()
}

func RandomIntStr() string {
	return strconv.FormatInt(rand.Int63(), 10)
}
