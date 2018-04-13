package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5FromReader(reader io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}

	if closer, is := reader.(io.Closer); is {
		closer.Close()
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Md5FromString(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}
