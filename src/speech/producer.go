package speech

import (
	"encoding/json"
	"httpserv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wserv"
)

func produce(auth string, message wserv.Message) {
	if auth == "" {
		return
	}

	message.Auth = ""
	message.Operation = "speech.consume"
	data, err := json.Marshal(message)
	if err != nil {
		return
	}

	push(auth, data)
	write(auth, message, data)
}

func push(auth string, data []byte) {
	for _, conn := range consumers[auth] {
		consume(conn, data)
	}
}

func write(auth string, message wserv.Message, data []byte) {
	path := getPath(auth, message.Type)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err)

		return
	}

	ioutil.WriteFile(path+message.Unique, data, 0644)
	if file, err := os.OpenFile(getUniques(auth), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err == nil {
		defer file.Close()
		file.WriteString(message.Type + ":" + message.Unique + "\n")
	}
}

func uniques(writer http.ResponseWriter, request *http.Request) int {
	auth := httpserv.GetParam(request, "auth", "")
	if auth == "" {
		return httpserv.Send404(writer)
	}

	return httpserv.ServeFile(writer, request, nil, getUniques(auth))
}
