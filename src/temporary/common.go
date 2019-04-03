package temporary

import (
	"log"
	"os"
	"path/filepath"
	"util"
)

type config struct {
	Root string
}

var cfg = config{"whtemp"}
var root = cfg.Root

func init() {
	if err := util.LoadConfig(&cfg, "temporary"); err != nil {
		return
	}

	root, _ = filepath.Abs(cfg.Root)
	os.MkdirAll(root, os.ModePerm)
	root = root + "/"
	cfg.Root = util.FormatPath("/" + cfg.Root + "/")
	log.Printf("temporary config: %+v\n", cfg)
}

func absolute(uri string) string {
	return util.FormatPath(root + uri)
}
