package memory

import (
	"log"
	"time"
	"util"
)

type config struct {
	Deadline int64
}

var cfg = config{1800}
var times = make(map[string]int64)
var bytes = make(map[string][]byte)

func init() {
	if err := util.LoadConfig(&cfg, "memory"); err != nil {
		return
	}

	log.Printf("memory config:deadline=%d\n", cfg.Deadline)

	go func() {
		for {
			time.Sleep(time.Second)
			clear()
		}
	}()
}

func update(unique string) {
	times[unique] = time.Now().Unix()
}
