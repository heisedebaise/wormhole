package util

import "math/rand"

var chars = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

// RandomString 随即字符串。
func RandomString(size int) string {
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = chars[rand.Intn(36)]
	}

	return string(bytes)
}
