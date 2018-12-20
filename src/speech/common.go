package speech

import (
	"log"
	"os"
	"util"
)

type config struct {
	Timeout  string
	nTimeout int64
}

var cfg = config{"5", 0}
var root = "speech/"

func init() {
	if err := util.LoadConfig(&cfg, "speech"); err != nil {
		return
	}
	cfg.nTimeout = int64(util.ParseTime(cfg.Timeout))
	log.Printf("speech config: %+v\n", cfg)
}

func createTime(auth string) int64 {
	if info, err := os.Stat(getPath(auth, "")); err == nil {
		return info.ModTime().Unix()
	}

	return -1
}

func modifyTime(auth string) int64 {
	if info, err := os.Stat(getUniques(auth)); err == nil {
		return info.ModTime().Unix()
	}

	return -1
}

func getUniques(auth string) string {
	return getPath(auth, "") + "uniques"
}

func getOutline(auth string) string {
	return getPath(auth, "") + "outline"
}

func getPath(auth string, t string) string {
	path := root + auth + "/"
	if t != "" {
		path += t + "/"
	}

	return path
}
