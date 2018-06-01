package util

import (
	"log"
	"regexp"
)

type whitelist struct {
	IP    []string `json:"Ip"`
	RegEx []string
}

var whitelistCfg whitelist
var hasWhitelistIP bool
var hasWhitelistRegEx bool

func init() {
	LoadConfig(&whitelistCfg, "whitelist")
	hasWhitelistIP = len(whitelistCfg.IP) > 0
	hasWhitelistRegEx = len(whitelistCfg.RegEx) > 0
	log.Printf("white list: %q\n", whitelistCfg)
}

// InWhiteList 校验是否存在于白名单内。
func InWhiteList(ip string) bool {
	if hasWhitelistIP {
		for _, str := range whitelistCfg.IP {
			if ip == str {
				return true
			}
		}
	}

	if hasWhitelistRegEx {
		for _, str := range whitelistCfg.RegEx {
			if match, err := regexp.MatchString(str, ip); err == nil && match {
				return true
			}
		}
	}

	return false
}
