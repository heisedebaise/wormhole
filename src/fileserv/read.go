package fileserv

import (
	"log"
	"net/http"
	"os"
	"protocol"
)

func read(writer http.ResponseWriter, request *http.Request, uri string) {
	uri = uri[len(cfg.Root):]
	path := absolute(uri)
	if info, err := os.Stat(path); err != nil || info.IsDir() {
		protocol.Send404(writer)
		log.Printf("not exists or read dir %s %q\n", path, err)
	} else {
		http.ServeFile(writer, request, path)
	}
}
