package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(oriString string) (encryptString string) {
	h := md5.New()
	// 加密盐
	h.Write([]byte("$12*&sdc"))
	// 转成16进制的字符串
	return hex.EncodeToString(h.Sum([]byte(oriString)))
}
