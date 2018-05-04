package main

import (
	"imgserv"
	"net/http"
	"protocol"
	"strings"
	"util"
)

type config struct {
	Listen string
}

var cfg config = config{":8192"}

func handler(writer http.ResponseWriter, request *http.Request) {
	uri := request.RequestURI
	if strings.HasPrefix(uri, imgserv.Root()) {
		imgserv.Handler(writer, request, uri)
	}
}

func main() {
	util.LoadConfig(&cfg, "http")
	protocol.Http(cfg.Listen, "/", handler)
}
