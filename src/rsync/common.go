package rsync

import (
	"util"

	"github.com/google/uuid"
)

type config struct {
	Listen    string
	Nodes     []string
	Argot     string
	ReConnect int
}

var fileFlag = byte(1)
var memoryFlag = byte(2)
var cfg = config{":2048", []string{"127.0.0.1:2048"}, "wormhome rsync argot", 5}
var id string

func init() {
	util.LoadConfig(&cfg, "rsync")
	id = uuid.New().String()
}
