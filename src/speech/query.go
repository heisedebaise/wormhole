package speech

import (
	"auth"
	"bufio"
	"bytes"
	"httpserv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

	file, err := os.Open(getUniques(consumer))
	if err != nil {
		return httpserv.SendFailure(writer, httpserv.Failure{Code: 2104, Message: "Uniques[" + consumer + "]不存在！"})
	}

	defer file.Close()
	t := httpserv.GetParam(request, "type", "")
	start := httpserv.GetParam(request, "start", "")
	end := httpserv.GetParam(request, "end", "")
	var buffer bytes.Buffer
	buffer.WriteString("[")
	first := true
	comma := []byte(",")
	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		line := scanner.Text()
		indexOf := strings.Index(line, ":")
		if indexOf == -1 || (t != "" && line[:indexOf] != t) {
			continue
		}

		unique := line[indexOf+1:]
		if start != "" && start > unique {
			continue
		}

		if end != "" && end < unique {
			break
		}

		if data, err := ioutil.ReadFile(getPath(consumer, t) + unique); err == nil {
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
