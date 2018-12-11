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
	Auth      string `json:"auth"`
	Operation string `json:"operation"`
	Unique    string `json:"unique"`
	Type      string `json:"type"`
	Content   string `json:"content"`
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
	log.Printf("websocket config: %+v\n", cfg)
}
