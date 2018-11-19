package wserv

import (
	"log"
	"util"

	"github.com/gorilla/websocket"
)

type config struct {
	Root string
}

// Message 消息格式。
type Message struct {
	Auth      string
	Operation string
	Unique    string
	Content   string
}

var cfg = config{"whws"}
var upgrader = websocket.Upgrader{}

func init() {
	if err := util.LoadConfig(&cfg, "websocket"); err != nil {
		return
	}

	cfg.Root = util.FormatPath("/" + cfg.Root)
	log.Printf("websocket config:root=%s\n", cfg.Root)
}
