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
	data, err := json.Marshal(failure)
	if err == nil {
		writer.Write(data)
	} else {
		writer.Write([]byte(err.Error()))
	}

	return Send200(writer)
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
