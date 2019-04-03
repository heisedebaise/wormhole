package httpserv

import (
	"crypto/tls"
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

	if cfg.SSL.Listen != "" && len(cfg.SSL.Certs) > 0 {
		httpsChan := make(chan int)
		go func() {
			tlsConfig := &tls.Config{}
			for _, cert := range cfg.SSL.Certs {
				if x509KeyPair, err := tls.LoadX509KeyPair(cert.Crt, cert.Key); err == nil {
					tlsConfig.Certificates = append(tlsConfig.Certificates, x509KeyPair)
				} else {
					log.Printf("fail to load cert %v %v\n", cert, err)
				}
			}
			if len(tlsConfig.Certificates) == 0 {
				log.Printf("cert is empty %v\n", cfg)

				return
			}

			log.Printf("https listening on %v\n", cfg.SSL)
			// tlsConfig.BuildNameToCertificate()
			// server := http.Server{
			// 	Addr:      cfg.SSL.Listen,
			// 	Handler:   nil,
			// 	TLSConfig: tlsConfig,
			// }
			// if err := server.ListenAndServeTLS("", ""); err != nil {
			if _, err := tls.Listen("tcp", cfg.SSL.Listen, tlsConfig); err != nil {
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
	setCors(writer, request)
	log.Println(request.URL.Port())
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
	}
	log.Printf("%d: uri=%s;remote=%s;time=%fms\n", code, uri, GetIP(request), float64((time.Now().UnixNano()-now))/1000000)
}

func setCors(writer http.ResponseWriter, request *http.Request) {
	origin := GetHeader(request, "Origin")
	if origin == "" {
		origin = "*"
	}
	if len(cfg.Cors.Origin) == 0 || (!contains(cfg.Cors.Origin, "*") && !contains(cfg.Cors.Origin, origin)) {
		return
	}

	SetHeader(writer, "Access-Control-Allow-Origin", origin)
	SetHeader(writer, "Access-Control-Allow-Methods", cfg.Cors.Methods)
	SetHeader(writer, "Access-Control-Allow-Headers", cfg.Cors.Headers)
	SetHeader(writer, "Access-Control-Allow-Credentials", "true")
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}

	return false
}

// Handler 添加处理器。
func Handler(root string, handler func(writer http.ResponseWriter, request *http.Request, uri string) int) {
	handlers[root] = handler
	log.Printf("bind http handler: %s\n", root)
}
