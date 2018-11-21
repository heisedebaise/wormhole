package wserv

import (
	"log"
	"net/http"
)

// Root 获取URI前缀。
func Root() string {
	return cfg.Root
}

// Handler 处理WS(S)请求。
func Handler(writer http.ResponseWriter, request *http.Request, uri string) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Printf("upgrade websocket %s failure %q\n", uri, err)

		return
	}

	defer conn.Close()
	read(conn)
}
