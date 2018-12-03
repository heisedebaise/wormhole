package auth

import (
	"log"
	"util"
)

type config struct {
	Root string
}

var cfg = config{"auth"}

func init() {
	if err := util.LoadConfig(&cfg, "auth"); err != nil {
		return
	}

	cfg.Root = util.FormatPath("/" + cfg.Root + "/")
	log.Printf("auth config:root=%s\n", cfg.Root)
}
