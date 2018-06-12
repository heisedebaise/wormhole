package protocol

import "util"

type httpcfg struct {
	Listen string
	RealIP string
}

var cfg = httpcfg{":8192", ""}

func init() {
	util.LoadConfig(&cfg, "http")
}
