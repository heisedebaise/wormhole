package wserv

import (
	"auth"
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

		producer := auth.GetProducer(msg.Auth)
		consumer := auth.GetConsumer(msg.Auth)
		if producer == "" && consumer == "" {
			break
		}

		register(conn, consumer)
		produce(producer, msg)
	}
}
