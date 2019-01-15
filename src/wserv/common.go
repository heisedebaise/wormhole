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

// Message 消息。
type Message struct {
	Auth      string `json:"auth"`
	Operation string `json:"operation"`
	Unique    string `json:"unique"`
	Type      string `json:"type"`
	State     string `json:"state"`
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
