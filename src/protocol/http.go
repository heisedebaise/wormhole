package protocol

import (
	"log"
	"net/http"
)

// HTTP 启动HTTP(S)服务。
func HTTP(path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		if cfg.Cors {
			writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		handler(writer, request)
	})

	log.Printf("http listening on %s\n", cfg.Listen)
	if err := http.ListenAndServe(cfg.Listen, nil); err != nil {
		log.Fatalln(err)
	}
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
