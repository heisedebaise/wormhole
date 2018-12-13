package speech

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
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
	produceTimes[auth] = time.Now().Unix()
}

func push(auth string, data []byte) {
	for _, conn := range consumers[auth] {
		consume(conn, data)
	}
}

func write(auth string, message wserv.Message, data []byte) {
	path := getPath(auth, message)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err)

		return
	}

	ioutil.WriteFile(path+message.Unique, data, 0644)
}
