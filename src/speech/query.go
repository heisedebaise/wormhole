package speech

import (
	"auth"
	"bytes"
	"httpserv"
	"io/ioutil"
	"net/http"
)

func outline(writer http.ResponseWriter, request *http.Request) int {
	auth := httpserv.GetParam(request, "auth", "")
	if auth == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2105, Message: "Auth不允许为空！"})
	}

	return httpserv.ServeFile(writer, request, nil, getOutline(auth))
}

func uniques(writer http.ResponseWriter, request *http.Request) int {
	auth := httpserv.GetParam(request, "auth", "")
	if auth == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2105, Message: "Auth不允许为空！"})
	}

	return httpserv.ServeFile(writer, request, nil, getUniques(auth))
}

func track(writer http.ResponseWriter, request *http.Request) int {
	ticket := httpserv.GetParam(request, "ticket", "")
	if ticket == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2103, Message: "Ticket不允许为空！"})
	}

	consumer := auth.GetConsumer(ticket)
	if consumer == "" {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2104, Message: "Ticket[" + ticket + "]认证失败！"})
	}

	path := getPath(consumer, httpserv.GetParam(request, "type", ""))
	if infos, err := ioutil.ReadDir(path); err == nil {
		start := httpserv.GetParam(request, "start", "")
		end := httpserv.GetParam(request, "end", "")
		var buffer bytes.Buffer
		buffer.WriteString("[")
		first := true
		comma := []byte(",")
		for _, info := range infos {
			name := info.Name()
			if start != "" && start > name {
				continue
			}

			if end != "" && end < name {
				break
			}

			if data, err := ioutil.ReadFile(path + name); err == nil {
				if first {
					first = false
				} else {
					buffer.Write(comma)
				}
				buffer.Write(data)
			}
		}
		buffer.WriteString("]")
		buffer.WriteTo(writer)

		return 200
	}

	return httpserv.Send404(writer)
}
