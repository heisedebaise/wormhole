package main

import (
	"fileserv"
	"httpserv"
	"imgserv"
	"log"
	"net/http"
	"rsync"
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
		code = httpserv.Send404(writer)
	}
	log.Printf("http%d: uri=%s;remote=%s;time=%fms\n", code, uri, httpserv.GetIP(request), float64((time.Now().UnixNano()-now))/1000000)
}

func main() {
	rsync.Listen()

	httpserv.HTTP("/", handler)
}
