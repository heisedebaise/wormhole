package speech

import (
	"auth"
	"httpserv"
	"io/ioutil"
	"net/http"
	"util"
)

func save(writer http.ResponseWriter, request *http.Request) int {
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	id := httpserv.GetParam(request, "id", "")
	data := httpserv.GetParam(request, "data", "")
	if id == "" || data == "" {
		return httpserv.Send404(writer)
	}

	ioutil.WriteFile(getData(id), []byte(data), 0644)
	writer.Write([]byte("success"))

	return 200
}

func data(writer http.ResponseWriter, request *http.Request) int {
	ticket := httpserv.GetParam(request, "ticket", "")
	if ticket == "" {
		return httpserv.Send404(writer)
	}

	token := auth.GetProducer(ticket)
	if token == "" {
		token = auth.GetConsumer(ticket)
	}
	if token == "" {
		return httpserv.Send404(writer)
	}

	return httpserv.ServeFile(writer, request, nil, getData(token))
}

func getData(auth string) string {
	return getPath(auth, "") + "data"
}
