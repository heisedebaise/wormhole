package rsync

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var storages = make(map[byte]func(unique string, message []byte))

// Storage 设置存储处理器。
func Storage(flag byte, storage func(unique string, message []byte)) {
	storages[flag] = storage
}

func saveFile(unique string, message []byte) {
	path, err := filepath.Abs(unique[1:])
	if err != nil {
		return
	}

	if err = os.MkdirAll(path[:strings.LastIndex(path, "/")], 0755); err != nil {
		return
	}

	ioutil.WriteFile(path, message, 0755)
}
