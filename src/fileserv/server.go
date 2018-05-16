package fileserv

import (
	"net/http"
	"protocol"
)

func Root() string {
	return cfg.Root
}

func Handler(writer http.ResponseWriter, request *http.Request, uri string) {
	if uri == cfg.Save {
		protocol.Upload(writer, request, maxSize, root, cfg.Root)
	} else {
		read(writer, request, uri)
	}
}
