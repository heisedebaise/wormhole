package protocol

import (
	"util"
)

type httpcfg struct {
	Listen string
	RealIP string
	Cors   bool
}

var cfg = httpcfg{":8192", "", false}

func init() {
	util.LoadConfig(&cfg, "http")
}
