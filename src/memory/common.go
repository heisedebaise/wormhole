package memory

import (
	"log"
	"rsync"
	"time"
	"util"
)

type config struct {
	Deadline int64
}

var cfg = config{1800}
var bytes = make(map[string][]byte)
var updates = make(map[string]int64)
var deadlines = make(map[string]int64)

func init() {
	if err := util.LoadConfig(&cfg, "memory"); err != nil {
		return
	}

	log.Printf("memory config: %+v\n", cfg)

	rsync.Storage(rsync.MemoryFlag, sync)
	go func() {
		for {
			time.Sleep(time.Second)
			clear()
		}
	}()
}

func update(unique string, sync bool) {
	time := time.Now().Unix() + cfg.Deadline
	updates[unique] = time
	if sync {
		rsync.SendBytes(rsync.MemoryFlag, unique, util.Int64ToBytes(time))
	}
}
