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

	consumers[consumer] = append(consumers[consumer], conn)
}

func produce(producer string, msg message) {
	if producer == "" {
		return
	}

	push(producer, msg.Unique, msg.Content)
}

func push(auth string, unique string, content string) {
	msg := message{}
	msg.Unique = unique
	msg.Operation = "speech.consume"
	msg.Content = content
	go func() {
		for _, conn := range consumers[auth] {
			if err := conn.WriteJSON(msg); err != nil {
				// consumers[auth] = append(consumers[auth][:index], consumers[auth][index+1:])
				conn.Close()
				log.Printf("send to websocket consumer failure %q !\n", err)
			}
		}
	}()
}
