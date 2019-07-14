package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Crypto(password string) string {
	data := []byte(password)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
