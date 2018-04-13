package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadConfig(config interface{}, name string) (err error) {
	data, err := ioutil.ReadFile("conf/" + name + ".json")
	if err != nil {
		log.Printf("Fail to load config file [conf/%s.json] %s \n", name, err)

		return
	}

	err = json.Unmarshal(data, config)

	return
}
