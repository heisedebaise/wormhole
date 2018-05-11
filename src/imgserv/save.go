package imgserv

import (
	"io/ioutil"
	"net/http"
	"os"
	"protocol"
	"strings"
)

func save(writer http.ResponseWriter, request *http.Request) {
	path, name, upload := protocol.Upload(writer, request, maxSize, root, cfg.Root)
	if upload {
		clean(path, name)
	}
}

func clean(path string, name string) {
	files, err := ioutil.ReadDir(absolute(path))
	if err != nil {
		return
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
