package rsync

import (
	"log"
	"util"

	"github.com/google/uuid"
)

type config struct {
	Listen    string
	Nodes     []string
	Argot     string
	ReConnect int
}

// FileFlag 文件标记。
var FileFlag = byte(1)

// MemoryFlag 内存标记。
var MemoryFlag = byte(2)

var cfg = config{":2048", []string{"127.0.0.1:2048"}, "wormhome rsync argot", 5}
var id string

func init() {
	if err := util.LoadConfig(&cfg, "rsync"); err != nil {
		return
	}

	log.Printf("rsync config: %+v\n", cfg)

	id = uuid.New().String()
	Storage(FileFlag, saveFile)
}
