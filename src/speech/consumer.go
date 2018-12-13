package speech

import (
	"io/ioutil"
	"strings"
	"wserv"

	"github.com/gorilla/websocket"
)

var consumers = make(map[string][]*websocket.Conn)
var consumerChans = make(map[*websocket.Conn]chan int)

func register(auth string, conn *websocket.Conn) {
	if auth == "" {
		return
	}

	consumers[auth] = append(consumers[auth], conn)
	consumerChans[conn] = make(chan int, 1)
}

func pull(auth string, message wserv.Message, conn *websocket.Conn) {
	if auth == "" {
		return
	}

	var start string
	var end string
	indexOf := strings.Index(message.Unique, ":")
	if indexOf == -1 {
		start = message.Unique[:indexOf]
		end = message.Unique[indexOf+1:]
	}
	path := getPath(auth, message)
	if files, err := ioutil.ReadDir(path); err == nil {
		consumerChans[conn] <- 0
		for _, file := range files {
			name := file.Name()
			if file.IsDir() || (start != "" && start > name) || (end != "" && end < name) {
				continue
			}

			if data, err := ioutil.ReadFile(path + name); err == nil {
				conn.WriteMessage(websocket.TextMessage, data)
			}
		}
		<-consumerChans[conn]
	}
}

func consume(conn *websocket.Conn, data []byte) {
	go func(conn *websocket.Conn, data []byte) {
		consumerChans[conn] <- 0
		conn.WriteMessage(websocket.TextMessage, data)
		<-consumerChans[conn]
	}(conn, data)
}
