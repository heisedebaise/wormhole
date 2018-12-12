package httpserv

import "net/http"

// Send404 发送404。
func Send404(writer http.ResponseWriter) int {
	return SendCode(writer, 404)
}

// SendCode 发送结果码。
func SendCode(writer http.ResponseWriter, code int) int {
	writer.WriteHeader(code)

	return code
}
