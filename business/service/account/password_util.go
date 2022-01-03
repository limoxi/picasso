package account

import (
	"crypto/sha1"
	"fmt"
)

// encodePassword 对raw string加密
func encodePassword(rawPwd string) string{
	h := sha1.New()
	h.Write([]byte(rawPwd))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// checkPassword 检查密码
func checkPassword(rawPwd, targetPwd string) bool{
	encodedPwd := encodePassword(rawPwd)
	return encodedPwd == targetPwd
}