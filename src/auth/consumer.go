package auth

import "memory"

func consumer(token string, ticket string) {
	memory.PutString("consumer:"+ticket, token, 0)
}

// GetConsumer 获取消费者认证信息。
func GetConsumer(ticket string) string {
	return memory.GetString("consumer:" + ticket)
}
