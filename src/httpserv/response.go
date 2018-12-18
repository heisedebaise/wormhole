package httpserv

import (
	"encoding/json"
	"net/http"
)

// Failure 失败信息。
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SendFailure 发送失败信息。
func SendFailure(writer http.ResponseWriter, failure Failure) int {
	if data, err := json.Marshal(failure); err == nil {
		writer.Write(data)

		return Send200(writer)
	}

	return Send404(writer)
}

// Send200 发送200。
func Send200(writer http.ResponseWriter) int {
	return SendCode(writer, 200)
}

// Send404 发送404。
func Send404(writer http.ResponseWriter) int {
	return SendCode(writer, 404)
}

// SendCode 发送结果码。
func SendCode(writer http.ResponseWriter, code int) int {
	writer.WriteHeader(code)

	return code
}
