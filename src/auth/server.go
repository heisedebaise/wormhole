package auth

import (
	"httpserv"
	"net/http"
)

func handle(writer http.ResponseWriter, request *http.Request, uri string) int {
	request.ParseForm()
	if code := httpserv.Auth(writer, request); code > 0 {
		return code
	}

	token := httpserv.GetParam(request, "token", "")
	if token == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 1001, Message: "Token不允许为空！"})
	}

	ticket := httpserv.GetParam(request, "ticket", "")
	if ticket == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 1002, Message: "Ticket不允许为空！"})
	}

	if uri == cfg.Root+"producer" {
		producer(token, ticket)
	} else if uri == cfg.Root+"consumer" {
		consumer(token, ticket)
	} else {
		return httpserv.Send404(writer)
	}

	return httpserv.SendSuccess(writer, nil)
}

// Serve 服务。
func Serve() {
	httpserv.Handler(cfg.Root, handle)
}
