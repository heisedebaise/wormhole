package protocol

import (
	"log"
	"net/http"
	"strings"
)

func Http(host string, path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(path, handler)
	log.Printf("listening on %s\n", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetParam(request *http.Request, name string, defaultValue string) string {
	if _, has := request.Form[name]; has {
		return request.Form[name][0]
	}

	return defaultValue
}

func GetIp(request *http.Request) string {
	ip := request.RemoteAddr
	ip = ip[0:strings.LastIndex(ip, ":")]
	if ip == "[::1]" {
		ip = "127.0.0.1"
	}

	return ip
}

func Send404(writer http.ResponseWriter) {
	SendCode(writer, 404)
}

func SendCode(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
}
