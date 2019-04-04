package imgserv

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	move()
}

func absolute(uri string) string {
	return util.FormatPath(root + uri)
}

func move() {
	go func() {
		for count := 0; count < 5; {
			file, err := os.Open(root)
			if err != nil {
				return
			}
			defer file.Close()

			infos, err := file.Readdir(1024)
			if err != nil {
				return
			}

			i := 0
			for _, info := range infos {
				if info.IsDir() {
					continue
				}

				name := info.Name()
				if strings.Index(name, ".") == 32 {
					source := absolute(name)
					target := absolute(util.Md5PathName(name))
					if os.Rename(source, target) == nil {
						i++
					} else {
						log.Printf("move %s to %s fail !\n", source, target)
					}
				}
			}
			log.Printf("move %d md5 file .\n", i)

			if i == 0 {
				count++
			}

			time.Sleep(time.Second)
		}
	}()
}
