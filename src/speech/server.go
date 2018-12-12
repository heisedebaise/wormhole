package speech

import (
	"auth"
	"wserv"

	"github.com/gorilla/websocket"
)

func handle(conn *websocket.Conn, message wserv.Message) {
	producer := auth.GetProducer(message.Auth)
	consumer := auth.GetConsumer(message.Auth)
	if producer == "" && consumer == "" {
		return
	}

	if message.Operation == "speech.consumer" {
		register(consumer, conn)
	} else if message.Operation == "speech.produce" {
		produce(producer, message)
	} else if message.Operation == "speech.pull" {
		pull(consumer, message, conn)
	} else if message.Operation == "speech.finish" {
		finish(consumer)
	}
}

// Serve 服务。
func Serve() {
	wserv.Handler("speech.", handle)
	auto()
}
