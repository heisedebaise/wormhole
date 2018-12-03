package auth

import "memory"

func consumer(auth string, unique string) {
	memory.PutString("consumer:"+unique, auth, 0)
}

// GetConsumer 获取消费者认证信息。
func GetConsumer(unique string) string {
	return memory.GetString("consumer:" + unique)
}
