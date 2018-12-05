package auth

import "memory"

func producer(token string, ticket string) {
	memory.PutString("producer:"+ticket, token, 0)
}

// GetProducer 获取生产者认证信息。
func GetProducer(ticket string) string {
	return memory.GetString("producer:" + ticket)
}
