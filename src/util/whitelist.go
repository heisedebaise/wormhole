package util

import (
	"log"
	"regexp"
)

type whitelist struct {
	Ip    []string
	RegEx []string
}

var whitelistCfg whitelist
var hasWhitelistIp bool
var hasWhitelistRegEx bool

func init() {
	LoadConfig(&whitelistCfg, "whitelist")
	hasWhitelistIp = len(whitelistCfg.Ip) > 0
	hasWhitelistRegEx = len(whitelistCfg.RegEx) > 0
	log.Printf("white list: %q\n", whitelistCfg)
}

func InWhiteList(ip string) bool {
	if hasWhitelistIp {
		for _, str := range whitelistCfg.Ip {
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
