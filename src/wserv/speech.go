package wserv

import (
	"auth"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var consumers = make(map[string][]*websocket.Conn)

func speech(conn *websocket.Conn, msg message) {
	if !strings.HasPrefix(msg.Operation, "speech.") {
		return
	}

	producer := auth.GetProducer(msg.Auth)
	consumer := auth.GetConsumer(msg.Auth)
	if producer == "" && consumer == "" {
		return
	}

	if msg.Operation == "speech.consumer" {
		register(consumer, conn)
	} else if msg.Operation == "speech.produce" {
		produce(producer, msg)
	} else if msg.Operation == "speech.pull" {
		pull(consumer, msg)
	}
}

func register(auth string, conn *websocket.Conn) {
	if auth == "" {
		return
	}

	consumers[auth] = append(consumers[auth], conn)
}

func produce(auth string, msg message) {
	if auth == "" {
		return
	}

	msg.Auth = ""
	msg.Operation = "speech.consume"
	push(auth, msg)
	write(auth, msg)
}

func pull(auth string, msg message) {
	if auth == "" {
		return
	}

	path := getPath(auth, msg)
	if path == "" {
		return
	}

	var start string
	var end string
	indexOf := strings.Index(msg.Unique, ":")
	if indexOf == -1 {
		start = msg.Unique[:indexOf]
		end = msg.Unique[indexOf+1:]
	}
	if files, err := ioutil.ReadDir(path); err == nil {
		for _, file := range files {
			name := file.Name()
			if file.IsDir() || (start != "" && start > name) || (end != "" && end < name) {
				continue
			}

			if data, err := ioutil.ReadFile(path + name); err == nil {
				var m message
				if json.Unmarshal(data, &m) == nil {
					push(auth, m)
				}
			}
		}
	}
}

func push(auth string, msg message) {
	// go func() {
	for _, conn := range consumers[auth] {
		if err := conn.WriteJSON(msg); err != nil {
			conn.Close()
			log.Printf("send to websocket consumer failure %q !\n", err)
		}
	}
	// }()
}

func write(auth string, msg message) {
	path := getPath(auth, msg)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err)

		return
	}

	if bytes, err := json.Marshal(msg); err == nil {
		ioutil.WriteFile(path+msg.Unique, bytes, 0644)
	}
}

func getPath(auth string, msg message) string {
	path := "speech/" + auth + "/"
	if msg.Type == "" {
		path += "type/"
	} else {
		path += msg.Type + "/"
	}

	return path
}
