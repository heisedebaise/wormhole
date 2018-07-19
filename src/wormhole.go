package main

import (
	"fileserv"
	"imgserv"
	"log"
	"net/http"
	"protocol"
	"strings"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	now := time.Now().UnixNano()
	uri := request.URL.Path
	code := 0
	if strings.HasPrefix(uri, imgserv.Root()) {
		code = imgserv.Handler(writer, request, uri)
	} else if strings.HasPrefix(uri, fileserv.Root()) {
		code = fileserv.Handler(writer, request, uri)
	} else {
		code = protocol.Send404(writer)
	}
	log.Printf("http%d: uri=%s;remote=%s;time=%fms\n", code, uri, protocol.GetIP(request), float64((time.Now().UnixNano()-now))/1000000)
}

func main() {
	protocol.HTTP("/", handler)
}
