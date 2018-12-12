package wserv

import (
	"httpserv"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var handlers = make(map[string]func(conn *websocket.Conn, message Message))

func handle(writer http.ResponseWriter, request *http.Request, uri string) int {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("upgrade websocket %s failure %q\n", uri, err)

		return 400
	}

	for {
		message := Message{}
		if err := conn.ReadJSON(&message); err != nil {
			log.Printf("read message from websocket failure %q\n", err)

			break
		}

		for prefix, handler := range handlers {
			if strings.HasPrefix(message.Operation, prefix) {
				handler(conn, message)

				break
			}
		}
	}

	return 200
}

// Handler 添加处理器。
func Handler(prefix string, handler func(conn *websocket.Conn, message Message)) {
	handlers[prefix] = handler
	log.Printf("bind web socket handler: %s\n", prefix)
}

// Serve 服务。
func Serve() {
	httpserv.Handler(cfg.URI, handle)
}
