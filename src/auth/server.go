package auth

import (
	"httpserv"
	"net/http"
	"util"
)

type parameter struct {
	Token  string
	Ticket string
}

func handle(writer http.ResponseWriter, request *http.Request, uri string) int {
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	var json = parameter{}
	if httpserv.GetJSON(request, &json) != nil {
		return httpserv.Send404(writer)
	}

	if json.Token == "" {
		return httpserv.Send404(writer)
	}

	if json.Ticket == "" {
		return httpserv.Send404(writer)
	}

	if uri == cfg.Root+"producer" {
		producer(json.Token, json.Ticket)
	} else if uri == cfg.Root+"consumer" {
		consumer(json.Token, json.Ticket)
	} else {
		return httpserv.Send404(writer)
	}

	writer.Write([]byte("success"))

	return 200
}

// Serve 服务。
func Serve() {
	httpserv.Handler(cfg.Root, handle)
}
