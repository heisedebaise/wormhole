package imgserv

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

var cfg = config{"whimg", "save", "10M"}
var root = cfg.Root
var maxSize int64 = 10 << 20

func init() {
	if err := util.LoadConfig(&cfg, "image"); err != nil {
		return
	}

	root, _ = filepath.Abs(cfg.Root)
	os.MkdirAll(root, os.ModePerm)
	root = root + "/"
	cfg.Root = util.FormatPath("/" + cfg.Root + "/")
	cfg.Save = util.FormatPath(cfg.Root + cfg.Save)
	maxSize = util.ByteSize(cfg.MaxSize)
	log.Printf("image config: %+v\n", cfg)
}

func absolute(uri string) string {
	return util.FormatPath(root + uri)
}
