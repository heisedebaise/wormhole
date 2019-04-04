package httpserv

import (
	"log"
	"strings"
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

func absolute(root string, path string, name string) string {
	length := strings.LastIndex(name, ".")
	if length == -1 {
		length = len(name)
	}

	abs := root + path
	i := 0
	for ; i < length; i += 2 {
		abs += "/" + name[i:i+2]
	}
	abs += name[i:]
	log.Println(util.FormatPath(abs))

	return util.FormatPath(root + path + "/" + name)
}

func md5PathName(name string) string {
	pathName := ""
	i := 0
	for ; i < 32; i += 2 {
		pathName += "/" + name[i:i+2]
	}
	if len(name) > 32 {
		pathName += name[i:]
	}

	return pathName
}
