package util

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfig(config interface{}, name string) (err error) {
	data, err := ioutil.ReadFile("conf/" + name + ".json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, config)

	return
}
