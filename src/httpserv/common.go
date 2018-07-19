package httpserv

import (
	"util"
)

type config struct {
	Listen string
	RealIP string
	Cors   bool
}

var cfg = config{":8192", "", false}

func init() {
	util.LoadConfig(&cfg, "http")
}
