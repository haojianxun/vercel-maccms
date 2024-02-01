package helpers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

const (
	base64Table = "bbUGUNiuy54851asaGYUGJBiugubnjkhu78584215jhbgbkjbuy+/"
)

func Md5(password string) string {
	str := password
	// 计算MD5哈希值
	hash := md5.Sum([]byte(str))
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}

func Base64Encode(src []byte) []byte { //编码
	return []byte(base64.NewEncoding(base64Table).EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) { //解码
	return base64.NewEncoding(base64Table).DecodeString(string(src))
}
