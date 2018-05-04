package imgserv

import (
	"net/http"
)

func Root() string {
	return cfg.Root
}

func Handler(writer http.ResponseWriter, request *http.Request, uri string) {
	if uri == cfg.Save {
		save(writer, request)
	} else {
		read(writer, request, uri)
	}
}
