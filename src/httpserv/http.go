package httpserv

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

	httpChan := make(chan int)
	go func() {
		log.Printf("http listening on %s\n", cfg.Listen)
		if err := http.ListenAndServe(cfg.Listen, nil); err != nil {
			log.Fatalln(err)
			httpChan <- -1
		} else {
			httpChan <- 1
		}
	}()

	if cfg.SSL.Listen != "" && cfg.SSL.Crt != "" && cfg.SSL.Key != "" {
		httpsChan := make(chan int)
		go func() {
			log.Printf("https listening on %s\n", cfg.SSL.Listen)
			if err := http.ListenAndServeTLS(cfg.SSL.Listen, cfg.SSL.Crt, cfg.SSL.Key, nil); err != nil {
				log.Fatalln(err)
				httpsChan <- -1
			} else {
				httpsChan <- 1
			}
		}()
		<-httpsChan
	}

	<-httpChan
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
