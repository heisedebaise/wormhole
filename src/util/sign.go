package util

import (
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type config struct {
	Name   string
	Secret string
}

var once sync.Once
var cfgs []config

func CheckSign(form url.Values) bool {
	if _, has := form["sign"]; !has {
		return false
	}

	if _, has := form["sign-time"]; !has {
		return false
	}

	signTime, err := strconv.Atoi(form["sign-time"][0])
	if err != nil || math.Abs(float64(time.Now().Unix()-int64(signTime/1000))) > 10 {
		return false
	}

	var keys []string
	for key, _ := range form {
		if key == "sign" {
			continue
		}

		keys = append(keys, key)
	}
	sort.Strings(keys)

	var str string
	for _, key := range keys {
		str += key + "=" + strings.Join(form[key], ",") + "&"
	}
	str += getSecret(form["sign-name"][0])

	return Md5FromString(str) == form["sign"][0]
}

func getSecret(name string) string {
	once.Do(func() {
		LoadConfig(cfgs, "sign")
	})

	for _, cfg := range cfgs {
		if cfg.Name == name {
			return cfg.Secret
		}
	}

	return ""
}
