package wserv

import (
	"log"
	"net/http"
)

// URI 获取URI。
func URI() string {
	return cfg.URI
}

// Handler 处理WS(S)请求。
func Handler(writer http.ResponseWriter, request *http.Request, uri string) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("upgrade websocket %s failure %q\n", uri, err)

		return
	}

	read(conn)
}
