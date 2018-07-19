package fileserv

import (
	"log"
	"net/http"
	"os"
	"protocol"
)

func read(writer http.ResponseWriter, request *http.Request, uri string) int {
	uri = uri[len(cfg.Root):]
	path := absolute(uri)
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		log.Printf("not exists or read dir %s %q\n", path, err)

		return protocol.Send404(writer)
	}

	return protocol.ServeFile(writer, request, info, path)
}
