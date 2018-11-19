package wserv

import (
	"log"

	"github.com/gorilla/websocket"
)

func read(conn *websocket.Conn) {
	for {
		message := Message{}
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("read message from websocket failure %q\n", err)

			break
		}
		log.Println(message)
		conn.WriteJSON(message)
	}
}
