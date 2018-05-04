package imgserv

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"util"
)

type config struct {
	Root    string
	Save    string
	MaxSize string
}

var cfg config = config{"whimg", "save", "10M"}
var root = cfg.Root
var maxSize int64 = 10 << 20

func init() {
	if err := util.LoadConfig(&cfg, "image"); err != nil {
		return
	}

	root, _ = filepath.Abs(cfg.Root)
	os.MkdirAll(root, os.ModePerm)
	root = root + "/"
	cfg.Root = strings.Replace("/"+cfg.Root+"/", "//", "/", -1)
	cfg.Save = strings.Replace(cfg.Root+cfg.Save, "//", "/", -1)
	maxSize = util.ByteSize(cfg.MaxSize)
	log.Printf("image config:root=%s;save=%s;absolute path=%s;max size=%d\n", cfg.Root, cfg.Save, root, maxSize)
}

func clean(path string, name string) {
	files, err := ioutil.ReadDir(absolute(path))
	if err != nil {
		return
	}

	if path != "" {
		name = name[strings.LastIndex(name, "/")+1 : len(name)]
	}

	names := strings.Split(name, ".")
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ns := strings.Split(file.Name(), ".")
		length := len(ns)
		if length > 2 && ns[0] == names[0] && ns[length-1] == names[1] {
			os.Remove(absolute(path + "/" + file.Name()))
		}
	}
}

func absolute(uri string) string {
	return strings.Replace(root+uri, "//", "/", -1)
}
