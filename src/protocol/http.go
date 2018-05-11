package protocol

import (
	"log"
	"net/http"
)

func Http(host string, path string, handler func(writer http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(path, handler)
	log.Printf("listening on %s\n", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Send404(writer http.ResponseWriter) {
	SendCode(writer, 404)
}

func SendCode(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
}
