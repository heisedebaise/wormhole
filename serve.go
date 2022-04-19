package wormhole

import (
	"encoding/json"
	"log"
	"os"
)

func Serve(config string) {
	bs, err := os.ReadFile(config)
	if err != nil {
		log.Printf("load %s err %v\n", config, err)

		return
	}

	cfg := struct {
		TCP  map[string]string `json:"tcp"`
		HTTP []httpcfg         `json:"http"`
	}{}
	if err = json.Unmarshal(bs, &cfg); err != nil {
		log.Printf("unmarshal %s err %v\n", config, err)

		return
	}
	log.Printf("load %s %v\n", config, cfg)

	tcp(cfg.TCP)
	https(cfg.HTTP)

	stat()
}
