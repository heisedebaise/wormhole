package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5FromReader 计算MD5值。
func Md5FromReader(reader io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Md5FromString 计算MD5值。
func Md5FromString(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}
