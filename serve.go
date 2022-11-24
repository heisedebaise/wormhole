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
		HTTP map[string]string `json:"http"`
	}{}
	if err = json.Unmarshal(bs, &cfg); err != nil {
		log.Printf("unmarshal %s err %v\n", config, err)

		return
	}
	log.Printf("load %s %v\n", config, cfg)

	tcp(cfg.TCP)
	serveHTTP(cfg.HTTP)

	stat()
}
