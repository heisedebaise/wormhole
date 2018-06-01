package fileserv

import (
	"log"
	"os"
	"path/filepath"
	"util"
)

type config struct {
	Root    string
	Save    string
	MaxSize string
}

var cfg = config{"whfile", "save", "10M"}
var root = cfg.Root
var maxSize int64 = 10 << 20

func init() {
	if err := util.LoadConfig(&cfg, "file"); err != nil {
		return
	}

	root, _ = filepath.Abs(cfg.Root)
	os.MkdirAll(root, os.ModePerm)
	root = root + "/"
	cfg.Root = util.FormatPath("/" + cfg.Root + "/")
	cfg.Save = util.FormatPath(cfg.Root + cfg.Save)
	maxSize = util.ByteSize(cfg.MaxSize)
	log.Printf("file config:root=%s;save=%s;absolute path=%s;max uploadable size=%d=%s\n", cfg.Root, cfg.Save, root, maxSize, cfg.MaxSize)
}

func absolute(uri string) string {
	return util.FormatPath(root + uri)
}
