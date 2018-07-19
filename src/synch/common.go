package synch

import (
	"util"

	"github.com/google/uuid"
)

type config struct {
	Listen string
	Nodes  []string
	Argot  string
}

var cfg = config{":2048", []string{"127.0.0.1:2048"}, "wormhome synch argot"}
var id string

func init() {
	util.LoadConfig(&cfg, "synch")
	id = uuid.New().String()
}
