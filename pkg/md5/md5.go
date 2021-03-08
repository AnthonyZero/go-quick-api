package md5

import (
	cryptoMD5 "crypto/md5"
	"encoding/hex"
)

var _ MD5 = (*md5)(nil)

//MD5 摘要
type MD5 interface {
	i()
	// Encrypt 加密返回32位
	Encrypt(encryptStr string) string
}

//New 方法 返回MD5类型
func New() MD5 {
	return &md5{}
}

type md5 struct{}

func (m *md5) i() {}

func (m *md5) Encrypt(encryptStr string) string {
	s := cryptoMD5.New()
	s.Write([]byte(encryptStr))
	return hex.EncodeToString(s.Sum(nil))
}
