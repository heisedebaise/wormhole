package main

import (
	"auth"
	"fileserv"
	"httpserv"
	"imgserv"
	"log"
	"net/http"
	"rsync"
	"strings"
	"time"
	"wserv"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	now := time.Now().UnixNano()
	uri := request.RequestURI
	schema := "http/s"
	code := 0
	if strings.HasPrefix(uri, imgserv.Root()) {
		code = imgserv.Handler(writer, request, uri)
	} else if strings.HasPrefix(uri, fileserv.Root()) {
		code = fileserv.Handler(writer, request, uri)
	} else if strings.HasPrefix(uri, auth.Root()) {
		code = auth.Handler(writer, request, uri)
	} else if uri == wserv.URI() {
		schema = "ws/s"
		wserv.Handler(writer, request, uri)
	} else {
		code = httpserv.Send404(writer)
	}
	log.Printf("%s %d: uri=%s;remote=%s;time=%fms\n", schema, code, uri, httpserv.GetIP(request), float64((time.Now().UnixNano()-now))/1000000)
}

func main() {
	rsync.Listen()

	httpserv.HTTP("/", handler)
}
