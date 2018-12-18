package speech

import (
	"auth"
	"httpserv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func save(writer http.ResponseWriter, request *http.Request) int {
	if code := httpserv.Auth(writer, request); code > 0 {
		return code
	}

	id := httpserv.GetParam(request, "id", "")
	if id == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2101, Message: "ID不允许为空！"})
	}

	data := httpserv.GetParam(request, "data", "")
	if id == "" || data == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2102, Message: "Data不允许为空！"})
	}

	path := getData(id)
	os.MkdirAll(path[strings.LastIndex(path, "/"):], os.ModePerm)
	ioutil.WriteFile(path, []byte(data), 0644)
	httpserv.SendSuccess(writer, nil)

	return 200
}

func data(writer http.ResponseWriter, request *http.Request) int {
	ticket := httpserv.GetParam(request, "ticket", "")
	if ticket == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2103, Message: "Ticket不允许为空！"})
	}

	token := auth.GetProducer(ticket)
	if token == "" {
		token = auth.GetConsumer(ticket)
	}
	if token == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2104, Message: "Ticket[" + ticket + "]认证失败！"})
	}

	return httpserv.ServeFile(writer, request, nil, getData(token))
}

func getData(auth string) string {
	return getPath(auth, "") + "data"
}
