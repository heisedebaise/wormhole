package wserv

import (
	"log"

	"github.com/gorilla/websocket"
)

func read(conn *websocket.Conn) {
	for {
		msg := message{}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("read message from websocket failure %q\n", err)

			break
		}

		speech(conn, msg)
	}
}
