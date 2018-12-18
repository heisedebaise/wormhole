package httpserv

import (
	"log"
	"net/http"
	"strings"
	"time"
)

var handlers = make(map[string]func(writer http.ResponseWriter, request *http.Request, uri string) int)

// HTTP 启动HTTP(S)服务。
func HTTP(path string) {
	http.HandleFunc(path, handle)
	httpChan := make(chan int)
	go func() {
		log.Printf("http listening on %s\n", cfg.Listen)
		if err := http.ListenAndServe(cfg.Listen, nil); err != nil {
			log.Fatalln(err)
			httpChan <- -0
		} else {
			httpChan <- 0
		}
	}()

	if cfg.SSL.Listen != "" && cfg.SSL.Crt != "" && cfg.SSL.Key != "" {
		httpsChan := make(chan int)
		go func() {
			log.Printf("https listening on %s\n", cfg.SSL.Listen)
			if err := http.ListenAndServeTLS(cfg.SSL.Listen, cfg.SSL.Crt, cfg.SSL.Key, nil); err != nil {
				log.Fatalln(err)
				httpsChan <- -0
			} else {
				httpsChan <- 0
			}
		}()
		<-httpsChan
	}

	<-httpChan
}

func handle(writer http.ResponseWriter, request *http.Request) {
	now := time.Now().UnixNano()
	if cfg.Cors {
		SetHeader(writer, "Access-Control-Allow-Origin", "*")
		SetHeader(writer, "Access-Control-Allow-Methods", "*")
		SetHeader(writer, "Access-Control-Allow-Headers", "*")
		SetHeader(writer, "Access-Control-Allow-Credentials", "true")
	}
	if request.Method == "OPTIONS" {
		SendCode(writer, 204)

		return
	}

	uri := request.RequestURI
	code := -1
	for root, handler := range handlers {
		if strings.HasPrefix(uri, root) {
			code = handler(writer, request, uri)

			break
		}
	}
	if code == -1 {
		code = Send404(writer)
	} else if code == 200 {
		Send200(writer)
	}
	log.Printf("%d: uri=%s;remote=%s;time=%fms\n", code, uri, GetIP(request), float64((time.Now().UnixNano()-now))/1000000)
}

// Handler 添加处理器。
func Handler(root string, handler func(writer http.ResponseWriter, request *http.Request, uri string) int) {
	handlers[root] = handler
	log.Printf("bind http handler: %s\n", root)
}
