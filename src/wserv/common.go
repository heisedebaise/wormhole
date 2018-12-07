package wserv

import (
	"log"
	"net/http"
	"util"

	"github.com/gorilla/websocket"
)

type config struct {
	URI string
}

type message struct {
	auth      string
	operation string
	unique    string
	content   string
}

var cfg = config{"whws"}
var upgrader = websocket.Upgrader{}

func init() {
	if err := util.LoadConfig(&cfg, "websocket"); err != nil {
		return
	}

	upgrader.CheckOrigin = func(request *http.Request) bool {
		return true
	}

	cfg.URI = util.FormatPath("/" + cfg.URI)
	log.Printf("websocket config:root=%s\n", cfg.URI)
}
