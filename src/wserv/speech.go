package wserv

import (
	"auth"
	"log"
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
		register(conn, consumer)
	} else if msg.Operation == "speech.produce" {
		produce(producer, msg)
	}
}

func register(conn *websocket.Conn, consumer string) {
	if consumer == "" {
		return
	}

	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("web socket %d %s close.\n", code, text)

		return nil
	})

	consumers[consumer] = append(consumers[consumer], conn)
}

func produce(auth string, msg message) {
	if auth == "" {
		return
	}

	push(auth, msg)
}

func push(auth string, msg message) {
	m := message{}
	m.Operation = "speech.consume"
	m.Unique = msg.Unique
	m.Type = msg.Type
	m.Content = msg.Content
	// go func() {
	for _, conn := range consumers[auth] {
		if err := conn.WriteJSON(m); err != nil {
			conn.Close()
			log.Printf("send to websocket consumer failure %q !\n", err)
		}
	}
	// }()
}

func write(msg message){
	
}
