package httpserv

import (
	"log"
	"util"
)

type config struct {
	Listen string
	RealIP string
	Cors   cors
	SSL    ssl
}

type cors struct {
	Origin  []string
	Methods string
	Headers string
}

type ssl struct {
	Listen string
	Certs  []cert
}

type cert struct {
	Crt string
	Key string
}

var cfg = config{":8192", "", cors{[]string{}, "GET,POST", ""}, ssl{":8193", []cert{{"conf/ssl/wormhole.crt", "conf/ssl/wormhole.key"}}}}

func init() {
	util.LoadConfig(&cfg, "http")

	log.Printf("http config: %+v\n", cfg)
}
