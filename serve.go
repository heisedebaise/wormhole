package wormhole

import (
	"encoding/json"
	"os"
)

func Serve(config string) {
	bs, err := os.ReadFile(config)
	if err != nil {
		Log("load %s err %v", config, err)

		return
	}

	cfg := struct {
		TCP     map[string]string `json:"tcp"`
		HTTP    map[string]string `json:"http"`
		Replace map[string]string `json:"replace"`
		Capture map[string]string `json:"capture"`
	}{}
	if err = json.Unmarshal(bs, &cfg); err != nil {
		Log("unmarshal %s err %v", config, err)

		return
	}
	Log("load %s %v", config, cfg)

	tcp(cfg.TCP)
	serveHTTP(cfg.HTTP, cfg.Replace, cfg.Capture)

	stat()
}
