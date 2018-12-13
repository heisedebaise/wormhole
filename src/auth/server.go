package auth

import (
	"httpserv"
	"net/http"
	"util"
)

func handle(writer http.ResponseWriter, request *http.Request, uri string) int {
	request.ParseForm()
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	token := httpserv.GetParam(request, "token", "")
	if token == "" {
		return httpserv.Send404(writer)
	}

	ticket := httpserv.GetParam(request, "ticket", "")
	if ticket == "" {
		return httpserv.Send404(writer)
	}

	if uri == cfg.Root+"producer" {
		producer(token, ticket)
	} else if uri == cfg.Root+"consumer" {
		consumer(token, ticket)
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
