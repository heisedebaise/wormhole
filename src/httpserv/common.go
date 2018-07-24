package httpserv

import (
	"util"
)

type config struct {
	Listen string
	RealIP string
	Cors   bool
	SSL    ssl
}

type ssl struct {
	Listen string
	Crt    string
	Key    string
}

var cfg = config{":8192", "", false, ssl{":8443", "conf/ssl/wormhole.crt", "conf/ssl/wormhole.key"}}

func init() {
	util.LoadConfig(&cfg, "http")
}
