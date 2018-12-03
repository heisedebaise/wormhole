package auth

import "memory"

func producer(auth string, unique string) {
	memory.PutString("producer:"+unique, auth, 0)
}

// GetProducer 获取生产者认证信息。
func GetProducer(unique string) string {
	return memory.GetString("producer:" + unique)
}
