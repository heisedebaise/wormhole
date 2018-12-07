package wserv

import (
	"auth"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

var consumers = make(map[string][]*websocket.Conn)

func speech(conn *websocket.Conn, msg message) {
	if !strings.HasPrefix(msg.operation, "speech.") {
		return
	}

	producer := auth.GetProducer(msg.auth)
	consumer := auth.GetConsumer(msg.auth)
	if producer == "" && consumer == "" {
		return
	}

	if msg.operation == "speech.consumer" {
		register(conn, consumer)
	} else if msg.operation == "speech.produce" {
		produce(producer, msg)
	}
}

func register(conn *websocket.Conn, consumer string) {
	if consumer == "" {
		return
	}

	log.Printf("register:%s\n", consumer)

	consumers[consumer] = append(consumers[consumer], conn)
}

func produce(producer string, msg message) {
	if producer == "" {
		return
	}

	log.Printf("produce:%s\n", producer)

	push(producer, msg.unique, msg.content)
}

func push(auth string, unique string, content string) {
	msg := message{}
	msg.unique = unique
	msg.operation = "speech.consume"
	msg.content = content
	go func() {
		for _, conn := range consumers[auth] {
			log.Printf("push:%s\n", auth)
			if err := conn.WriteJSON(msg); err != nil {
				delete(consumers, auth)
				conn.Close()
				log.Printf("send to websocket consumer failure %q !\n", err)
			}
		}
	}()
}
