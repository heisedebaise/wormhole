package httpserv

import (
	"encoding/json"
	"net/http"
)

// Success 成功信息。
type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Failure 失败信息。
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SendSuccess 发送成功信息。
func SendSuccess(writer http.ResponseWriter, data interface{}) int {
	if data == nil {
		writer.Write([]byte("{code:0}"))
	} else {
		WriteJSON(writer, Success{Code: 0, Data: data})
	}

	return Send200(writer)
}

// SendFailure 发送失败信息。
func SendFailure(writer http.ResponseWriter, failure Failure) int {
	WriteJSON(writer, failure)

	return Send200(writer)
}

// WriteJSON 输出JSON数据。
func WriteJSON(writer http.ResponseWriter, v interface{}) {
	if data, err := json.Marshal(v); err == nil {
		writer.Write(data)
	}
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
