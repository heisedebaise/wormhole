package protocol

import (
	"log"
	"net/http"
)

// HTTP 启动HTTP(S)服务。
func HTTP(host string, path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(path, handler)
	log.Printf("listening on %s\n", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Send404 发送404。
func Send404(writer http.ResponseWriter) {
	SendCode(writer, 404)
}

// SendCode 发送结果码。
func SendCode(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
}
