package speech

import (
	"log"
	"util"
)

type config struct {
	Timeout  string
	nTimeout int64
}

var cfg = config{"8h", 0}

func init() {
	if err := util.LoadConfig(&cfg, "speech"); err != nil {
		return
	}
	cfg.nTimeout = int64(util.ParseTime(cfg.Timeout))
	log.Printf("speech config: %+v\n", cfg)
}

func getPath(auth string, t string) string {
	path := "speech/" + auth + "/"
	if t != "" {
		path += t + "/"
	}

	return path
}
