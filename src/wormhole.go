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
	if strings.HasPrefix(uri, imgserv.Root()) {
		imgserv.Handler(writer, request, uri)
	} else if strings.HasPrefix(uri, fileserv.Root()) {
		fileserv.Handler(writer, request, uri)
	}
	log.Printf("uri=%s;remote=%s;nano-time=%d\n", uri, protocol.GetIP(request), time.Now().UnixNano()-now)
}

func main() {
	protocol.HTTP("/", handler)
}
