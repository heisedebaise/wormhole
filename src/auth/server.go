package auth

import (
	"httpserv"
	"log"
	"net/http"
	"util"
)

// Root 获取URI前缀。
func Root() string {
	return cfg.Root
}

type parameter struct {
	Auth   string
	Unique string
}

// Handler 处理HTTP(S)请求。
func Handler(writer http.ResponseWriter, request *http.Request, uri string) int {
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	var json = parameter{}
	if httpserv.GetJSON(request, &json) != nil {
		return httpserv.Send404(writer)
	}

	log.Println(json)
	if json.Auth == "" {
		return httpserv.Send404(writer)
	}

	if json.Unique == "" {
		return httpserv.Send404(writer)
	}

	if uri == cfg.Root+"producer" {
		producer(json.Auth, json.Unique)
	} else if uri == cfg.Root+"consumer" {
		consumer(json.Auth, json.Unique)
	} else {
		return httpserv.Send404(writer)
	}

	writer.Write([]byte("success"))

	return 200
}
