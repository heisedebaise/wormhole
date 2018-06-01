package util

import (
	"log"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type sign struct {
	Name   string
	Secret string
}

var signCfg []sign

func init() {
	LoadConfig(&signCfg, "sign")
}

// CheckSign 校验签名。
func CheckSign(form url.Values) bool {
	if _, has := form["sign"]; !has {
		log.Println("no sign parameter !")

		return false
	}

	if _, has := form["sign-time"]; !has {
		log.Println("no sign-time parameter !")

		return false
	}

	signTime, err := strconv.Atoi(form["sign-time"][0])
	if err != nil || math.Abs(float64(time.Now().Unix()-int64(signTime/1000))) > 10 {
		log.Println("sign-time parameter illegal !")

		return false
	}

	var keys []string
	for key := range form {
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

	signName := ""
	if _, has := form["sign-name"]; has {
		signName = form["sign-name"][0]
		log.Printf("use sign-name=%s.\n", signName)
	}
	str += getSecret(signName)
	if Md5FromString(str) != form["sign"][0] {
		log.Println("sign parameter illegal !")

		return false
	}

	return true
}

func getSecret(name string) string {
	for _, cfg := range signCfg {
		if cfg.Name == name {
			return cfg.Secret
		}
	}

	log.Printf("use default sign name for %s.\n", name)

	return ""
}
