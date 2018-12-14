package speech

import (
	"auth"
	"httpserv"
	"net/http"
	"util"
	"wserv"

	"github.com/gorilla/websocket"
)

func h(writer http.ResponseWriter, request *http.Request, uri string) int {
	request.ParseForm()
	if !util.InWhiteList(httpserv.GetIP(request)) && !util.CheckSign(request.Form) {
		return httpserv.Send404(writer)
	}

	switch uri {
	case "/whspeech/outline":
		return outline(writer, request)
	case "/whspeech/uniques":
		return uniques(writer, request)
	case "/whspeech/track":
		return track(writer, request)
	default:
		return httpserv.Send404(writer)
	}
}

func ws(conn *websocket.Conn, message wserv.Message) {
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
	httpserv.Handler("/whspeech/", h)
	wserv.Handler("speech.", ws)
	scan()
}
